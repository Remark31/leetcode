package leetcode


func minMTA(a int , b int ) int {
	if a > b {
		return b
	}
	return a
}

func findKth(nums1 []int, nums2 []int, start1 int, len1 int, start2 int, len2 int, k int) int {
	if len1 > len2{
		return findKth(nums2, nums1, start2, len2, start1, len1, k)
	}

	if len1 == 0 {
		return nums2[start2 + k - 1]
	}

	if (k == 1) {
		return min(nums1[start1], nums2[start2]);
	}
	
	p := minMTA(k / 2,len1)
	if p <= 0{
		p = 1
	}
	q := k - p

	if nums1[start1 + p - 1] == nums2[start2 + q - 1]  {
		return nums1[start1 + p - 1];
	} else if nums1[start1 + p - 1] > nums2[start2 + q - 1] {
		return findKth(nums1, nums2, start1, len1, start2 + q, len2 - q, k - q);
	} else {
		return findKth(nums1, nums2, start1 + p, len1 - p, start2, len2, k - p);
	}
	
}


func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	n1 := len(nums1)
	n2 := len(nums2)
	if (n1 + n2) % 2 == 0 {
		return float64(findKth(nums1, nums2, 0 , n1 , 0 , n2 , (n1+n2)/2) + findKth(nums1, nums2, 0 , n1 , 0 , n2 , (n1+n2)/2+1))/2
	}
	return float64(findKth(nums1, nums2, 0 , n1, 0 , n2 , (n1+n2)/2 + 1))

}
