import cp_const as c
import cp_check_poker_hand as m

"""
test 牌
"""
poker_deck = m.PokerDeck()
# poker_deck.shuffle_deck()
# deck = poker_deck.deck
# print(poker_deck.deck)
# print(m.is_royal_straight_flush(deck[0:5]))

"""
test 牌型 function 
"""

## is_five_of_a_kind
# tmp_cards = []
# tmp_cards.append(m.convert_to_value("黑桃", 7))
# tmp_cards.append(m.convert_to_value("黑桃", 7))
# tmp_cards.append(m.convert_to_value("愛心", 7))
# tmp_cards.append(m.convert_to_value("梅花", 7))
# tmp_cards.append(m.convert_to_value("方塊", 7))
# print(m.is_five_of_a_kind(tmp_cards))


## is_straight_flush
# tmp_cards = []
# tmp_cards.append(m.convert_to_value("黑桃", 7))
# tmp_cards.append(m.convert_to_value("黑桃", 8))
# tmp_cards.append(m.convert_to_value("黑桃", 9))
# tmp_cards.append(m.convert_to_value("黑桃", 10))
# tmp_cards.append(m.convert_to_value("黑桃", 11))
# print(m.is_straight_flush(tmp_cards))


# # is_five_of_gold_tiger
# tmp_cards = []
# tmp_cards.append(m.convert_to_value("黑桃", 6))
# tmp_cards.append(m.convert_to_value("黑桃", 6))
# tmp_cards.append(m.convert_to_value("梅花", 6))
# tmp_cards.append(m.convert_to_value("方塊", 6))
# tmp_cards.append(m.convert_to_value("黑桃", 11))
# gold = 6
# print(m.is_five_of_gold_tiger(tmp_cards, gold))


# # is_four_of_a_kind
# tmp_cards = []
# tmp_cards.append(m.convert_to_value("黑桃", 7))
# tmp_cards.append(m.convert_to_value("愛心", 7))
# tmp_cards.append(m.convert_to_value("方塊", 7))
# tmp_cards.append(m.convert_to_value("梅花", 7))
# tmp_cards.append(m.convert_to_value("愛心", 5))
# print(m.is_four_of_a_kind(tmp_cards))


## is_gold_full_house
# tmp_cards = []
# tmp_cards.append(m.convert_to_value("黑桃", 7))
# tmp_cards.append(m.convert_to_value("方塊", 7))
# tmp_cards.append(m.convert_to_value("愛心", 7))
# tmp_cards.append(m.convert_to_value("梅花", 5))
# tmp_cards.append(m.convert_to_value("方塊", 5))
# gold = 7
# print(m.is_gold_full_house(tmp_cards, gold))

# # is_full_house
# tmp_cards = []
# tmp_cards.append(m.convert_to_value("黑桃", 7))
# tmp_cards.append(m.convert_to_value("方塊", 7))
# tmp_cards.append(m.convert_to_value("愛心", 7))
# tmp_cards.append(m.convert_to_value("梅花", 5))
# tmp_cards.append(m.convert_to_value("方塊", 5))
# gold = 6
# print(m.is_full_house(tmp_cards))

# # is_pair_with_flush
# tmp_cards = []
# tmp_cards.append(m.convert_to_value("黑桃", 7))
# tmp_cards.append(m.convert_to_value("黑桃", 7))
# tmp_cards.append(m.convert_to_value("黑桃", 2))
# tmp_cards.append(m.convert_to_value("黑桃", 6))
# tmp_cards.append(m.convert_to_value("黑桃", 5))
# print(m.is_pair_with_flush(tmp_cards))


# # is_flush
# tmp_cards = []
# tmp_cards.append(m.convert_to_value("黑桃", 9))
# tmp_cards.append(m.convert_to_value("黑桃", 6))
# tmp_cards.append(m.convert_to_value("黑桃", 5))
# tmp_cards.append(m.convert_to_value("黑桃", 4))
# tmp_cards.append(m.convert_to_value("黑桃", 3))
# print(m.is_flush(tmp_cards))


## is_straight
# tmp_cards = []
# tmp_cards.append(m.convert_to_value("黑桃", 10))
# tmp_cards.append(m.convert_to_value("黑桃", 11))
# tmp_cards.append(m.convert_to_value("愛心", 12))
# tmp_cards.append(m.convert_to_value("梅花", 13))
# tmp_cards.append(m.convert_to_value("方塊", 1))
# print(m.is_straight(tmp_cards))


# # is_gold_three_of_a_kind
# tmp_cards = []
# tmp_cards.append(m.convert_to_value("黑桃", 10))
# tmp_cards.append(m.convert_to_value("黑桃", 11))
# tmp_cards.append(m.convert_to_value("愛心", 7))
# tmp_cards.append(m.convert_to_value("梅花", 7))
# tmp_cards.append(m.convert_to_value("方塊", 7))
# gold = 7
# print(m.is_gold_three_of_a_kind(tmp_cards, gold))


# # is_three_of_a_kind
# tmp_cards = []
# tmp_cards.append(m.convert_to_value("黑桃", 1))
# tmp_cards.append(m.convert_to_value("愛心", 4))
# tmp_cards.append(m.convert_to_value("愛心", 3))
# tmp_cards.append(m.convert_to_value("梅花", 3))
# tmp_cards.append(m.convert_to_value("方塊", 3))
# print(m.is_three_of_a_kind(tmp_cards))

# is_two_pairs
# tmp_cards = []
# tmp_cards.append(m.convert_to_value("黑桃", 1))
# tmp_cards.append(m.convert_to_value("愛心", 1))
# tmp_cards.append(m.convert_to_value("愛心", 4))
# tmp_cards.append(m.convert_to_value("梅花", 3))
# tmp_cards.append(m.convert_to_value("方塊", 3))
# print(m.is_two_pairs(tmp_cards))


# is_one_pair
# tmp_cards = []
# tmp_cards.append(m.convert_to_value("黑桃", 1))
# tmp_cards.append(m.convert_to_value("愛心", 2))
# tmp_cards.append(m.convert_to_value("愛心", 4))
# tmp_cards.append(m.convert_to_value("梅花", 3))
# tmp_cards.append(m.convert_to_value("方塊", 3))
# print(m.is_one_pair(tmp_cards))


# # is_high_card
# tmp_cards = []
# tmp_cards.append(m.convert_to_value("黑桃", 1))
# tmp_cards.append(m.convert_to_value("愛心", 2))
# tmp_cards.append(m.convert_to_value("愛心", 3))
# tmp_cards.append(m.convert_to_value("梅花", 6))
# tmp_cards.append(m.convert_to_value("方塊", 7))
# print(m.is_high_card(tmp_cards))
