## CSS-Grid
  1. CSS-Grid is a two-dimentional system.(row x column).
  2. Flexbox  is a simpler and one-dimentional.
	
**Css layout**
we used 
1. table	
2. floats positioning and inline-block

**Solution.**
1. Flexbox  (but it's one-dimentional.)
2. CSS-Grid (two-dimentional)
	
> Best solution:
   Flexbox and Grid actually work very well together.

### To get Start:
	step 1. Container should have display:grid
	step 2. Set grid-template-columns and grid-template-rows
	step 3. In place its child elements into the grid with grid-column and grid-row

### Terminology:		
**1. Grid Container:**
Direct parent of all the grid items.

ex. In this example container is the grid container.
 ```
		<div class="container">
			<div class="item item-1"></div>
			<div class="item item-2"></div>
			<div class="item item-3"></div>
		</div>
```
**2. Grid Item:**
The children (e.g. direct descendants) of the grid container.

ex. Here the item elements are grid items, but sub-item isn't.
```			
			<div class="container">
				<div class="item"></div> 
				<div class="item">
					<p class="sub-item"></p>
				</div>
				<div class="item"></div>
			</div>
```
**3. Grid Line:**
The dividing lines that make up the structure of the grid.
- vertical line (column grid line)
- horizontal line ( row grid line)

**4. Grid Track:**
The space between two adjacent grid lines. (You can think of them like the columns or rows of the grid).

**5. Grid Cell: (Single Unit)**
The space between two adjacent row and two adjacent column grid lines.

**6. Grid Area:**	
The total space surrounded by four grid lines.
    - A grid area may be comprised of any number of grid cells.

### Properties for the Grid Container
- display:  grid | inline-grid;
- grid-template-columns: \<track-size\> ... | \<line-name\> \<track-size\> ...;
  - Values:
    - \<track-size\> \- can be a length, a percentage, or a fraction of the free space in the grid (using the fr unit)
    - \<line-name\> \- an arbitrary name of your choosing
       - ex
```
	   .container {
  		grid-template-columns: 40px 50px auto 50px 40px;
  		grid-template-rows: 25% 100px auto;
	  }
```
- grid-template-rows: \<track-size\> ... | \<line-name\> \<track-size\> ...;
  - Values:
    - \<track-size\> \- can be a length, a percentage, or a fraction of the free space in the grid (using the fr unit)
    - \<line-name\> \- an arbitrary name of your choosing
    	-  you can choose to explicitly name the lines. Note the bracket syntax for the line names:
```
.container {
	grid-template-columns: [first] 40px [line2] 50px [line3] auto [col4-start] 50px [five] 40px [end];
	grid-template-rows: [row1-start] 25% [row1-end] 100px [third-line] auto [last-line];
     }
```

- grid-template-areas: "\<grid-area-name\> | . | none | ..."
	- Values:
	    - \<grid-area-name\> => the name of a grid area specified with grid-area
	    - \.  => a period signifies an empty grid cell
	    - none => no grid areas are defined 
```
.item-a {
  grid-area: header;
}
.item-b {
  grid-area: main;
}
.item-c {
  grid-area: sidebar;
}
.item-d {
  grid-area: footer;
}

.container {
  grid-template-columns: 50px 50px 50px 50px;
  grid-template-rows: auto;
  grid-template-areas: 
    "header header header header"
    "main main . sidebar"
    "footer footer footer footer";
}
```
	    
	    
- grid-template
- grid-column-gap
- grid-row-gap
- grid-gap
- justify-items
- align-items
- place-items
- justify-content
- align-content
- place-content
- grid-auto-columns
- grid-auto-rows
- grid-auto-flow
- grid
	
### Properties for the Grid Items
- grid-column-start
- grid-column-end
- grid-row-start
- grid-row-end
- grid-column
- grid-row
- grid-area
- justify-self
- align-self
- place-self    	 
