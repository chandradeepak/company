# company


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
