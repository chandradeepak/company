package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"strings"
)

const (
	quotes = "\""
	comma  = ","
)

var (
	errInputDataFormatWrong    = "wrong number of delimeters or input format wrong data:%s"
	errEmpIDEmpty              = "employee ID empty :%s"
	errDuplicateEmp            = "duplicate employee found Name:%s ID:%s "
	errCircularDependencyFound = "circular dependency found  emp ID: %s manager ID: %s"
	errEmployeeWithoutManager  = "more than one employee with out manager, Name: %s, ID: %s, Manager ID: %s"
	errManagerDoesNtExist      = "manager for this employee doesn't exist emp ID: %s manager ID: %s"
)

//since it is mentioned the company is small we are gong to go with in memory tree stucture and going
// to use map for quick reference. for large companies it would be easy to store data in database
type employee struct {
	ID        string
	Manager   *employee
	Employees []*employee
	ManagerID string
	Name      string
}

//prints the heirarchy of the employess
func printCompanyHeirarchy(e *employee, level int) {

	//for each level we need that many tabs
	for i := 0; i <= level; i++ {
		fmt.Print("\t")
	}
	//for each name put a new line
	fmt.Println(e.Name)
	l := level + 1
	for _, emp := range e.Employees {
		printCompanyHeirarchy(emp, l)
	}

}

func parseRecord(record string) (*employee, error) {

	data := strings.Split(record, comma)

	if len(data) != 3 {
		return nil, fmt.Errorf(errInputDataFormatWrong, data)
	}

	name := strings.Trim(strings.TrimSpace(data[0]), quotes)
	id := strings.Trim(strings.TrimSpace(data[1]), quotes)
	managerID := strings.Trim(strings.TrimSpace(data[2]), quotes)

	if id == "" {
		return nil, fmt.Errorf(errEmpIDEmpty, data)
	}

	//make sure your are trimmimg all spaces when putting in map
	e := &employee{Name: name, ID: id, ManagerID: managerID}
	return e, nil

}

func loadData(fileName string) (map[string]*employee, error) {
	employeeMap := make(map[string]*employee)
	file, err := os.Open(fileName)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	//read one line at a time
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		//read comma delimited data
		emp, err := parseRecord(scanner.Text())
		if err != nil {
			return nil, err
		}
		//update employee map with id and the employee
		de, ok := employeeMap[emp.ID]
		if ok {
			return nil, fmt.Errorf(errDuplicateEmp, de.Name, de.ID)
		}
		employeeMap[emp.ID] = emp

	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}
	return employeeMap, nil
}

// Link the employees by iterating through the map and assign each employees manager and his employees
// this linking can be further optimized when reading the data itself but for code readbility purpose we are
// iterating in seperately
func linkRelationShip(employeeMap map[string]*employee) (*employee, error) {
	var ceo *employee
	for _, emp := range employeeMap {

		//check if employee's manager exists. If he exists create the employee - manager relation ship
		// and add this employee as sub ordinate of the manager.
		m, ok := employeeMap[emp.ManagerID]
		if ok {
			emp.Manager = m
			m.Employees = append(m.Employees, emp)
			employeeMap[emp.ID] = emp
			employeeMap[m.ID] = m
			if m.ManagerID == emp.ID {
				return nil, fmt.Errorf(errCircularDependencyFound, emp.ID, m.ManagerID)
			}
		} else {
			// if we have got only one employee who doesn't have a manager. we are assuming data is given like that
			if ceo != nil {
				return nil, fmt.Errorf(errEmployeeWithoutManager, emp.Name, emp.ID, emp.ManagerID)
			}
			if emp.ManagerID != "" {
				return nil, fmt.Errorf(errManagerDoesNtExist, emp.ID, emp.ManagerID)
			}
			log.Println("can not find manager for employee, ceo is:", emp.Name, emp.ID, emp.ManagerID)
			ceo = emp
		}

	}
	return ceo, nil
}

func main() {

	dataFile := flag.String("data", "data.csv", "file name")
	flag.Parse()

	log.Println("data file is", *dataFile)

	employeeMap, err := loadData(*dataFile)
	if err != nil {
		log.Fatal(err)
	}

	ceo, err := linkRelationShip(employeeMap)
	if err != nil {
		log.Fatal(err)
	}
	if ceo == nil {
		log.Fatal("there is no ceo for this company")
	}

	//print the heirarchy of employees
	printCompanyHeirarchy(ceo, 0)
}
