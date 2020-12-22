package day22

import (
	"bytes"
	"fmt"
	"hash/maphash"
	"strconv"
	"strings"
)

func Part1(input string) (string, error) {
	return solve(input, 1)
}

func Part2(input string) (string, error) {
	return solve(input, 2)
}

type deck []uint8

func newDeck(cards []uint8) *deck {
	d := make(deck, len(cards))
	copy(d, cards)
	return &d
}

func (d *deck) Empty() bool {
	return len(*d) == 0
}

func (d *deck) PopFront() uint8 {
	v := (*d)[0]
	*d = (*d)[1:]
	return v
}

func (d *deck) PushBack(v uint8) {
	*d = append(*d, v)
}

func (d *deck) String() string {
	ss := make([]string, len(*d))
	for i, card := range *d {
		ss[i] = fmt.Sprint(card)
	}
	return strings.Join(ss, ",")
}

func parsePlayer(paragraph string) []uint8 {
	lines := strings.Split(paragraph, "\n")[1:] // skip the first line
	cards := make([]uint8, len(lines))
	for i, line := range lines {
		card, err := strconv.Atoi(line)
		if err != nil {
			panic(err)
		}
		cards[i] = uint8(card)
	}
	return cards
}

func combat(deck1, deck2 *deck) *deck {
	for !deck1.Empty() && !deck2.Empty() {
		card1 := deck1.PopFront()
		card2 := deck2.PopFront()
		if card1 > card2 {
			deck1.PushBack(card1)
			deck1.PushBack(card2)
		} else {
			deck2.PushBack(card2)
			deck2.PushBack(card1)
		}
	}
	if !deck1.Empty() {
		return deck1
	}
	return deck2
}

type result struct {
	d  *deck
	nr int
}

type cacheEntry struct {
	deck1, deck2 *deck
	res          result
}

type cache struct {
	m map[uint64][]cacheEntry
	h maphash.Hash
}

func newCache() *cache {
	return &cache{
		m: make(map[uint64][]cacheEntry),
	}
}

func (c *cache) hash(deck1, deck2 *deck) uint64 {
	c.h.Reset()
	c.h.Write(*deck1)
	c.h.WriteByte(0) // assumes that 0 does not appear in the input
	c.h.Write(*deck2)
	return c.h.Sum64()
}

func (c *cache) Get(deck1, deck2 *deck) (result, bool) {
	hash := c.hash(deck1, deck2)
	entries, ok := c.m[hash]
	if !ok {
		return result{}, false
	}
	for _, entry := range entries {
		if bytes.Equal(*entry.deck1, *deck1) && bytes.Equal(*entry.deck2, *deck2) {
			return entry.res, true
		}
	}
	return result{}, false
}

func (c *cache) Set(deck1, deck2 *deck, res result) {
	hash := c.hash(deck1, deck2)
	entry := cacheEntry{
		deck1: deck1,
		deck2: deck2,
		res:   res,
	}
	c.m[hash] = append(c.m[hash], entry)
}

func subGame(deck1, deck2 *deck, c *cache) (winner *deck) {
	if r, ok := c.Get(deck1, deck2); ok {
		if r.nr == 1 {
			*deck1 = *r.d
			*deck2 = nil
			return deck1
		}
		*deck2 = *r.d
		*deck1 = nil
		return deck2
	}
	defer func() {
		d := make(deck, len(*winner))
		copy(d, *winner)
		nr := 1
		if winner == deck2 {
			nr = 2
		}
		c.Set(deck1, deck2, result{
			d:  &d,
			nr: nr,
		})
	}()

	// memo holds the set of previously seen configurations in this
	// game. Its elements are computed using the memoKey function.
	gameCache := newCache()
	for !deck1.Empty() && !deck2.Empty() {
		if _, ok := gameCache.Get(deck1, deck2); ok {
			return deck1
		}
		gameCache.Set(deck1, deck2, result{}) // ugly hack to only store presence
		card1 := deck1.PopFront()
		card2 := deck2.PopFront()
		var (
			roundWinner   *deck
			first, second uint8
		)
		if uint8(len(*deck1)) >= card1 && uint8(len(*deck2)) >= card2 {
			subDeck1 := make(deck, card1)
			copy(subDeck1, (*deck1)[:card1])
			subDeck2 := make(deck, card2)
			copy(subDeck2, (*deck2)[:card2])
			subWinner := subGame(&subDeck1, &subDeck2, c)
			switch subWinner {
			case &subDeck1:
				roundWinner = deck1
				first = card1
				second = card2
			case &subDeck2:
				roundWinner = deck2
				first = card2
				second = card1
			}
		} else {
			if card1 > card2 {
				roundWinner = deck1
				first = card1
				second = card2
			} else {
				roundWinner = deck2
				first = card2
				second = card1
			}
		}
		roundWinner.PushBack(first)
		roundWinner.PushBack(second)
	}
	if !deck1.Empty() {
		return deck1
	}
	return deck2
}

func recursiveCombat(deck1, deck2 *deck) *deck {
	return subGame(deck1, deck2, newCache())
}

func solve(input string, part int) (string, error) {
	paragraphs := strings.Split(strings.TrimSpace(input), "\n\n")
	deck1 := newDeck(parsePlayer(paragraphs[0]))
	deck2 := newDeck(parsePlayer(paragraphs[1]))
	var winner *deck
	switch part {
	case 1:
		winner = combat(deck1, deck2)
	case 2:
		winner = recursiveCombat(deck1, deck2)
	}
	score := 0
	f := len(*winner)
	for _, card := range *winner {
		score += int(card) * f
		f--
	}
	return fmt.Sprint(score), nil
}
