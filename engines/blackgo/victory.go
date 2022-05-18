package engine

import (
	"blackgo/deck"
	dTypes "blackgo/deck/types"
	"blackgo/utils"
	"sort"
)

func checkBlackGo(d deck.Deck) bool {
	return highestValidCombination(d) == 21
}

func isOutOfPlay(d deck.Deck) bool {
	return highestValidCombination(d) > 21
}

/**
Return the highest valid sum of the deck
*/
func highestValidCombination(d deck.Deck) int {
	combinations := allDeckSumCombinations(d)
	sort.Ints(combinations)

	validCombinations, outbound := utils.Separate(combinations, func(t int) bool { return t <= 21 })

	if len(validCombinations) > 0 {
		return validCombinations[len(validCombinations)-1]
	}

	return outbound[len(outbound)-1]
}

func allDeckSumCombinations(d deck.Deck) []int {
	aces, others := utils.Separate(d, func(card dTypes.Card) bool {
		return card.Number == dTypes.CA
	})

	ace_variations := generateAceVariations(len(aces))
	others_sum := utils.Sum(utils.Map(others, func(t dTypes.Card) int { return MinInt(int(t.Number), 10) }), 0)

	return utils.Map(ace_variations, func(t int) int { return t + others_sum })
}

func MinInt(a, b int) int {
	if a > b {
		return b
	}
	return a
}

func generateAceVariations(totalCards int) []int {
	variations := make([]int, totalCards+1)

	for i := 0; i < totalCards+1; i++ {
		elements := createArrayOf(1, totalCards-i)
		remaining := createArrayOf(11, i)
		result := append(elements, remaining...)

		variations = append(variations, utils.Sum(result, 0))
	}

	return variations
}

func createArrayOf(element int, size int) []int {
	output := []int{}
	for i := 0; i < size; i++ {
		output = append(output, element)
	}
	return output
}
