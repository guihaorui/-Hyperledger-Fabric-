package main

import (
	"fmt"
	"math"
)

// 模拟 PriceConstants 结构
type PriceConstants struct {
	RatingThreshold       float64
	IncreaseRatePerPoint  float64
	DecreaseRatePerPoint  float64
	MaxChangeRate         float64
}

// 模拟 Asset 结构
type Asset struct {
	AvgRating             float64
	PreviousQualityScore  float64
	Price                 int
}

// 模拟 calculateDeltaQ 函数
func calculateDeltaQ(asset *Asset, constants *PriceConstants) float64 {
	Q_current := asset.AvgRating
	// Q_previous := asset.PreviousQualityScore // 不再使用

	// 计算当前评分与5分阈值的差异
	diff_from_threshold := Q_current - constants.RatingThreshold

	// 如果差异为0（即正好5分），价格不变
	if math.Abs(diff_from_threshold) < 1e-10 {
		return 0.0
	}

	// 根据差异方向使用不同的每分变化率
	var ratePerPoint float64
	if diff_from_threshold > 0 {
		ratePerPoint = constants.IncreaseRatePerPoint  // 每分增加0.6%
	} else {
		ratePerPoint = constants.DecreaseRatePerPoint  // 每分减少1.2%
	}

	// 计算基础变化率：差异 * 每分变化率
	base_change := diff_from_threshold * ratePerPoint

	// 应用最大变化率限制
	if base_change > constants.MaxChangeRate {
		base_change = constants.MaxChangeRate
	} else if base_change < -constants.MaxChangeRate {
		base_change = -constants.MaxChangeRate
	}

	return base_change
}

// 模拟 calculateNewPrice 函数（简化版，忽略vTrend和C）
func calculateNewPrice(oldPrice int, deltaQ float64) int {
	P_old := float64(oldPrice)
	P_new := P_old * (1 + deltaQ)  // 忽略vTrend和C，C=0
	// 确保价格不为负，并取整
	if P_new < 0 {
		P_new = 0
	}
	// 确保价格不为零（至少为1）
	if P_new < 1 {
		P_new = 1
	}
	return int(math.Round(P_new))
}

func main() {
	// 测试常量
	constants := &PriceConstants{
		RatingThreshold:       5.0,
		IncreaseRatePerPoint:  0.006,  // 0.6%
		DecreaseRatePerPoint:  0.012,  // 1.2%
		MaxChangeRate:         0.06,   // 6%
	}

	// 测试用例：原始价格100元
	basePrice := 100

	testCases := []struct{
		name      string
		rating    float64
		expectedDeltaQ float64
		expectedPrice float64
	}{
		// 五分不变
		{"5分 - 不变", 5.0, 0.0, 100.0},
		// 高于5分：增加
		{"6分 - +0.6%", 6.0, 0.006, 100.6},
		{"7分 - +1.2%", 7.0, 0.012, 101.2},
		{"8分 - +1.8%", 8.0, 0.018, 101.8},
		{"9分 - +2.4%", 9.0, 0.024, 102.4},
		{"10分 - +3.0%", 10.0, 0.03, 103.0},
		// 低于5分：减少（双倍惩罚）
		{"4分 - -1.2%", 4.0, -0.012, 98.8},
		{"3分 - -2.4%", 3.0, -0.024, 97.6},
		{"2分 - -3.6%", 2.0, -0.036, 96.4},
		{"1分 - -4.8%", 1.0, -0.048, 95.2},
		{"0分 - -6.0%", 0.0, -0.06, 94.0},
		// 边界测试
		{"5.5分 - +0.3%", 5.5, 0.003, 100.3},
		{"4.5分 - -0.6%", 4.5, -0.006, 99.4},
	}

	fmt.Println("=== 评分价格变化测试 ===")
	fmt.Printf("基础价格: %d元\n", basePrice)
	fmt.Println()

	allPassed := true
	for _, tc := range testCases {
		asset := &Asset{
			AvgRating: tc.rating,
			PreviousQualityScore: 5.0, // 假设之前是5分
			Price: basePrice,
		}

		deltaQ := calculateDeltaQ(asset, constants)
		newPrice := calculateNewPrice(basePrice, deltaQ)

		// 计算预期价格（四舍五入）
		expectedPrice := int(math.Round(tc.expectedPrice))

		// 验证
		deltaQPass := math.Abs(deltaQ - tc.expectedDeltaQ) < 1e-10
		pricePass := newPrice == expectedPrice

		status := "✓"
		if !deltaQPass || !pricePass {
			status = "✗"
			allPassed = false
		}

		fmt.Printf("%s %s: 评分%.1f → ΔQ=%.4f (预期%.4f)", status, tc.name, tc.rating, deltaQ, tc.expectedDeltaQ)
		if !deltaQPass {
			fmt.Printf(" [ΔQ错误: 实际%.4f != 预期%.4f]", deltaQ, tc.expectedDeltaQ)
		}
		fmt.Printf(" 价格=%d元 (预期%d元)", newPrice, expectedPrice)
		if !pricePass {
			fmt.Printf(" [价格错误: 实际%d != 预期%d]", newPrice, expectedPrice)
		}
		fmt.Println()
	}

	fmt.Println()
	if allPassed {
		fmt.Println("✅ 所有测试通过！")
	} else {
		fmt.Println("❌ 部分测试失败，需要检查逻辑")
	}
}