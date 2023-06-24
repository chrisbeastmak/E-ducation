--[[

-- The following code is an example of a Lua program that can be used to teach students about the basics of E-ducation.

-- This program will print "Hello, E-ducation!" when run.

print("Hello, E-ducation!")

-- This code will calculate the sum of three numbers:

num1 = 5
num2 = 7
num3 = 10

sum = num1 + num2 + num3
print("The sum of", num1, ",", num2, "and", num3, "is", sum)

-- This code will show a message depending on the value of a variable:

grade = 7

if grade == 10 then
    print("Excellent job!")
elseif grade >= 8 and grade < 10 then
    print("You did very well!")
elseif grade >= 5 and grade < 8 then
    print("You did ok.")
else
    print("You need to work harder.")
end

-- This code will display all odd numbers between 1 and 10:

for n=1, 10, 2 do 
    print(n)
end

-- This code will prompt the user to enter their name and then say hello:

name =  io.read()
print("Hello, " .. name .. "!")

-- This code will create a table containing three items:

fruit = { "apple", "banana", "pear" }

-- This code will loop through the items in the table and print them on separate lines:

for i=1, #fruit do
    print(fruit[i])
end

-- This code will define a function called addNumbers that takes two arguments and returns the sum of those two numbers:

function addNumbers(num1, num2)
    return num1 + num2
end

-- This code will use the addNumbers function to calculate the sum of 5 and 7:

sum = addNumbers(5, 7)
print("The sum of 5 and 7 is", sum)

--]]