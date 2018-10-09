import java.util.Arrays;

public class TwoSum {
    public static void main(String[] args) {
        Solution solution = new Solution();
        int nums[] = {2, 7, 11, 15};
        int answer[] = solution.twoSum(nums, 9);
        System.out.println(Arrays.toString(answer));
    }
}
