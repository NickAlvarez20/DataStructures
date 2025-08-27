//Intuition: In-place sorting array, using two pointers. Iterate array and add all nums that don't match value to newArr, return newArr

function removeElement(nums: number[], val: number): number {
  let j: number = 0;
  for (let i = 0; i < nums.length; i++) {
    if (nums[i] !== val) {
      nums[j] = nums[i];
      j++;
    }
  }
  return j;
}
