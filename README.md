# packed-my-bag
[I packed my bag](http://en.wikipedia.org/wiki/I_packed_my_bag) as a console game written in go

# Example
The > character indicates an input. The screen is cleared after each block. After putting in a new object, it waits 3 seconds before clearing.

```
Number Of Players?:
>2
```
```
Player 1 
Add New Object:
>Shirt
```
```
Player 2 
Remember Object 1:
>Shirt
```
```
Player 2 
Add New Object:
>Towel
```
```
Player 1 
Remember Object 1:
>Shirt
```
```
Player 1 
Remember Object 2:
>Towel
```
```
Player 1 
Add New Object:
>Glasses
```
```
Player 2 
Remember Object 1:
>Shirt
```
```
Player 2 
Remember Object 2:
>Glasses
```
```
Wrong, your said Glasses but Towel was correct.
Player 2 lost the game!

Content of the bag:
1: Shirt
2: Towel
3: Glasses
```
