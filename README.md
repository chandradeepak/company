# company

        
## Assummptions
-  Since it is given it is a small company, solution is developed in memory not with out sql  database
-  Only one employee exists with out manager and he is the ceo
-  Having more than one employee with out manager is going to stop the program
-  The hierarchy is printed always from ceo
-  The top structure always start with one single employee who is CEO
-  CEO always exists in the data
-  Having Circular dependency is going to fail the program
-  Incorrect data given is going to stop the program
-  Incorrect data include no employee ID and normal employee not having Manager
-  Incorrect data also include where a manager ID exists but that employee ID doesn't exist
-  Input is expected to be given in csv format with commas in between for the ease of parsing data
-  Data is assumed to be fitting in memory so the data structures used are in memory. we used a tree and a map
-  The length of names should not be too long for the ease of printing so that you can see clearly how the hierarchy data is printed


## Input Data
```bash
    Input  has to be given in csv format for easy parsing purpose.
    Alan, 	        100, 	150 
    Martin, 	    220, 	100 
    Jamie, 	        150, 	 
    Alex, 	        275, 	100 
    Steve, 	        400, 	150 
    David, 	        190, 	400
    Chandra,        401,   190
    Divya,          402,   190
    Karhtik,        403,    275
    Nithya,         404,    275
```


## How to run
```bash
  go build && ./company -data=data.csv

   Would print some thing like this
   Jamie
                Alan
                        Martin
                        Alex
                                Karhtik
                                Nithya
                Steve
                        David
                                Divya
                                Chandra
```


## Printing html
Running would make a data.html like below. which you can open to see the html table
```html

			<html>
			<head>
			<style>
			table, th, td {
			border: 1px solid black;
			}
			</style>
			</head>
			<body>
			<table style="width:100%">
					
					
					<tr>
					
					
						<td>Jamie</td>
					
					</tr>
					
					<tr>
					
					
						<td>	</td>
					
						<td>Alan</td>
					
					</tr>
					
					<tr>
					
					
						<td>	</td>
					
						<td>	</td>
					
						<td>Martin</td>
					
					</tr>
					
					<tr>
					
					
						<td>	</td>
					
						<td>	</td>
					
						<td>Alex</td>
					
					</tr>
					
					<tr>
					
					
						<td>	</td>
					
						<td>Steve</td>
					
					</tr>
					
					<tr>
					
					
						<td>	</td>
					
						<td>	</td>
					
						<td>David</td>
					
					</tr>
					
					<tr>
					
					
					</tr>
					
			</table>
			</body>
			</html>
			
```



## Testing
```bash
go test -v

=== RUN   Test_Parsing
=== RUN   Test_Parsing/quotes_in_record
=== RUN   Test_Parsing/more_than_3_delimeters
=== RUN   Test_Parsing/empty_ID
--- PASS: Test_Parsing (0.00s)
    --- PASS: Test_Parsing/quotes_in_record (0.00s)
    --- PASS: Test_Parsing/more_than_3_delimeters (0.00s)
    --- PASS: Test_Parsing/empty_ID (0.00s)
=== RUN   Test_LoadData
=== RUN   Test_LoadData/empty_ID
=== RUN   Test_LoadData/employee_ID_empty
=== RUN   Test_LoadData/wrong_data_format_ID
--- PASS: Test_LoadData (0.00s)
    --- PASS: Test_LoadData/empty_ID (0.00s)
    --- PASS: Test_LoadData/employee_ID_empty (0.00s)
    --- PASS: Test_LoadData/wrong_data_format_ID (0.00s)
=== RUN   Test_LinkRelationShip
=== RUN   Test_LinkRelationShip/no_ceo
=== RUN   Test_LinkRelationShip/circular_dependency
--- PASS: Test_LinkRelationShip (0.00s)
    --- PASS: Test_LinkRelationShip/no_ceo (0.00s)
    --- PASS: Test_LinkRelationShip/circular_dependency (0.00s)
PASS
ok  	github.devtools.predix.io/predix-data-fabric/company	1.342s
```