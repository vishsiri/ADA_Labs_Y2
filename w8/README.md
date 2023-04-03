//64050229
//Vish Siriwatana
//THIS CODE WRITE IN GOLANG

**The algorithm in this code aims to find all possible binary solutions to a problem and evaluate them to see if they satisfy a given condition.**

> **Here is the high-level description of the algorithm:**

- 1) Initialize the necessary variables, including an array of integers w, which represents the weights of a set of items, and two integers n
   and M, which represent the number of items and the maximum weight
   that can be carried, respectively.
- 2) Calculate the total number of binary solutions as 2^n, where n is the number of items.

- 3) Iterate over all possible binary solutions by performing a loop that starts from 1 and ends at 2^n - 1.

- 4) For each solution, convert the solution into binary representation and store it in an array x.

- 5) Calculate the weight of the items selected in the solution by using the formula sum = solutionX(ix,x,w), where ix is the current solution and x and w are the arrays representing the binary solution and weights of the items, respectively.

- 6) Print the binary representation of the solution and the corresponding weight.

- 7) If the weight is equal to the maximum weight M, mark the solution with an asterisk.

- 8) Repeat steps 4 to 7 for all possible solutions.

- 9) End the program.

  ----
  
***ฟังก์ชั่น solutionX** ใช้สำหรับแปลงเลขฐานสิบเป็นเลขฐานสองและเก็บไว้ในอาร์เรย์ b โดยการทำงานคือหาร X ด้วย 2 และเก็บเศษที่ได้ในตำแหน่ง i ของอาร์เรย์ b แล้วนำ X ไปหารด้วย 2 จนกว่า X จะเป็น 0 โดยที่ตัวแปร i จะเพิ่มขึ้นทุกครั้งที่ได้เก็บเลขฐานสองของเลขฐานสิบที่เหลืออยู่ใน X เรียบร้อยแล้ว*

*ฟังก์ชั่น Pow2n ใช้สำหรับคำนวณค่า 2 ยกกำลัง n โดยการใช้ operator ชื่อ **<< (bitwise left shift)** กับตัวเลข 1 โดยจะทำการ shift ตัวเลข 1 ไปทางซ้าย n บิต ทำให้เกิดเลขฐานสองที่มีจำนวน 0 ที่เหลือเป็นตัวเลข 1 จำนวน n ตัว*