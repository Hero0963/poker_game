import cp_const as c
import cp_check_poker_hand as m
import cp_calculate_combat as cal
import cp_poker_hand_combinations as cbn
from typing import List
import collections

# # has_five_of_a_kind
# tmp_cards = []
# tmp_cards.append(m.convert_to_value("黑桃", 7))
# tmp_cards.append(m.convert_to_value("黑桃", 7))
# tmp_cards.append(m.convert_to_value("愛心", 7))
# tmp_cards.append(m.convert_to_value("梅花", 7))
# tmp_cards.append(m.convert_to_value("方塊", 7))
# tmp_cards.append(m.convert_to_value("方塊", 1))
# tmp_cards.append(m.convert_to_value("方塊", 4))
# tmp_cards.append(m.convert_to_value("方塊", 6))
# tmp_cards.append(m.convert_to_value("方塊", 5))
# tmp_cards.append(m.convert_to_value("方塊", 8))
# print(cbn.has_five_of_a_kind(tmp_cards))


# # has_straight_flush
# tmp_cards = []
# tmp_cards.append(m.convert_to_value("黑桃", 7))
# tmp_cards.append(m.convert_to_value("黑桃", 8))
# tmp_cards.append(m.convert_to_value("黑桃", 9))
# tmp_cards.append(m.convert_to_value("黑桃", 10))
# tmp_cards.append(m.convert_to_value("黑桃", 11))
# tmp_cards.append(m.convert_to_value("方塊", 1))
# tmp_cards.append(m.convert_to_value("方塊", 4))
# tmp_cards.append(m.convert_to_value("方塊", 6))
# tmp_cards.append(m.convert_to_value("方塊", 5))
# tmp_cards.append(m.convert_to_value("方塊", 8))
# print(cbn.has_straight_flush(tmp_cards))

# # has_straight_flush
# tmp_cards = []
# tmp_cards.append(m.convert_to_value("黑桃", 2))
# tmp_cards.append(m.convert_to_value("黑桃", 3))
# tmp_cards.append(m.convert_to_value("黑桃", 4))
# tmp_cards.append(m.convert_to_value("黑桃", 5))
# tmp_cards.append(m.convert_to_value("黑桃", 1))
# tmp_cards.append(m.convert_to_value("方塊", 1))
# tmp_cards.append(m.convert_to_value("方塊", 4))
# tmp_cards.append(m.convert_to_value("方塊", 6))
# tmp_cards.append(m.convert_to_value("方塊", 5))
# tmp_cards.append(m.convert_to_value("方塊", 8))
# print(cbn.has_straight_flush(tmp_cards))


# # has_five_of_gold_tiger
# tmp_cards = []
# tmp_cards.append(m.convert_to_value("黑桃", 6))
# tmp_cards.append(m.convert_to_value("黑桃", 6))
# tmp_cards.append(m.convert_to_value("梅花", 6))
# tmp_cards.append(m.convert_to_value("方塊", 6))
# tmp_cards.append(m.convert_to_value("黑桃", 11))
# tmp_cards.append(m.convert_to_value("黑桃", 2))
# tmp_cards.append(m.convert_to_value("黑桃", 3))
# tmp_cards.append(m.convert_to_value("梅花", 5))
# tmp_cards.append(m.convert_to_value("方塊", 7))
# tmp_cards.append(m.convert_to_value("黑桃", 8))
# gold = 6
# print(cbn.has_five_of_gold_tiger(tmp_cards, gold))


# # has_four_of_a_kind
# tmp_cards = []
# tmp_cards.append(m.convert_to_value("黑桃", 7))
# tmp_cards.append(m.convert_to_value("愛心", 7))
# tmp_cards.append(m.convert_to_value("方塊", 7))
# tmp_cards.append(m.convert_to_value("梅花", 7))
# tmp_cards.append(m.convert_to_value("愛心", 5))
# tmp_cards.append(m.convert_to_value("黑桃", 5))
# tmp_cards.append(m.convert_to_value("愛心", 6))
# tmp_cards.append(m.convert_to_value("方塊", 2))
# tmp_cards.append(m.convert_to_value("梅花", 3))
# tmp_cards.append(m.convert_to_value("愛心", 4))
# print(cbn.has_four_of_a_kind(tmp_cards))


# # has_gold_full_house
# tmp_cards = []
# tmp_cards.append(m.convert_to_value("黑桃", 7))
# tmp_cards.append(m.convert_to_value("方塊", 7))
# tmp_cards.append(m.convert_to_value("愛心", 7))
# tmp_cards.append(m.convert_to_value("梅花", 5))
# tmp_cards.append(m.convert_to_value("方塊", 5))
# tmp_cards.append(m.convert_to_value("黑桃", 8))
# tmp_cards.append(m.convert_to_value("方塊", 9))
# tmp_cards.append(m.convert_to_value("愛心", 10))
# tmp_cards.append(m.convert_to_value("梅花", 3))
# tmp_cards.append(m.convert_to_value("方塊", 4))
# gold = 7
# print(cbn.has_gold_full_house(tmp_cards, gold))


# # has_full_house
# tmp_cards = []
# tmp_cards.append(m.convert_to_value("黑桃", 7))
# tmp_cards.append(m.convert_to_value("方塊", 7))
# tmp_cards.append(m.convert_to_value("愛心", 7))
# tmp_cards.append(m.convert_to_value("梅花", 5))
# tmp_cards.append(m.convert_to_value("方塊", 5))
# tmp_cards.append(m.convert_to_value("黑桃", 8))
# tmp_cards.append(m.convert_to_value("方塊", 9))
# tmp_cards.append(m.convert_to_value("愛心", 10))
# tmp_cards.append(m.convert_to_value("梅花", 3))
# tmp_cards.append(m.convert_to_value("方塊", 4))
# gold = 6
# print(cbn.has_full_house(tmp_cards))

# # has_pair_with_flush
# tmp_cards = []
# tmp_cards.append(m.convert_to_value("黑桃", 7))
# tmp_cards.append(m.convert_to_value("黑桃", 7))
# tmp_cards.append(m.convert_to_value("愛心", 7))
# tmp_cards.append(m.convert_to_value("愛心", 6))
# tmp_cards.append(m.convert_to_value("愛心", 5))
# tmp_cards.append(m.convert_to_value("黑桃", 8))
# tmp_cards.append(m.convert_to_value("方塊", 9))
# tmp_cards.append(m.convert_to_value("愛心", 10))
# tmp_cards.append(m.convert_to_value("梅花", 3))
# tmp_cards.append(m.convert_to_value("方塊", 4))
# print(cbn.has_pair_with_flush(tmp_cards))
#
#
# # has_pair_with_flush
# tmp_cards = []
# tmp_cards.append(m.convert_to_value("黑桃", 7))
# tmp_cards.append(m.convert_to_value("黑桃", 7))
# tmp_cards.append(m.convert_to_value("愛心", 7))
# tmp_cards.append(m.convert_to_value("梅花", 7))
# tmp_cards.append(m.convert_to_value("方塊", 8))
# tmp_cards.append(m.convert_to_value("方塊", 9))
# tmp_cards.append(m.convert_to_value("方塊", 4))
# tmp_cards.append(m.convert_to_value("愛心", 10))
# tmp_cards.append(m.convert_to_value("愛心", 5))
# print(cbn.has_pair_with_flush(tmp_cards))

# # has_flush
# tmp_cards = []
# tmp_cards.append(m.convert_to_value("黑桃", 9))
# tmp_cards.append(m.convert_to_value("黑桃", 6))
# tmp_cards.append(m.convert_to_value("黑桃", 5))
# tmp_cards.append(m.convert_to_value("黑桃", 4))
# tmp_cards.append(m.convert_to_value("黑桃", 3))
# tmp_cards.append(m.convert_to_value("方塊", 8))
# tmp_cards.append(m.convert_to_value("方塊", 9))
# tmp_cards.append(m.convert_to_value("方塊", 4))
# tmp_cards.append(m.convert_to_value("愛心", 10))
# tmp_cards.append(m.convert_to_value("黑桃", 2))
# print(cbn.has_flush(tmp_cards))

# # has_straight
# tmp_cards = []
# tmp_cards.append(m.convert_to_value("黑桃", 10))
# tmp_cards.append(m.convert_to_value("黑桃", 11))
# tmp_cards.append(m.convert_to_value("愛心", 12))
# tmp_cards.append(m.convert_to_value("梅花", 13))
# tmp_cards.append(m.convert_to_value("方塊", 1))
# tmp_cards.append(m.convert_to_value("方塊", 7))
# tmp_cards.append(m.convert_to_value("方塊", 9))
# tmp_cards.append(m.convert_to_value("方塊", 4))
# tmp_cards.append(m.convert_to_value("愛心", 10))
# tmp_cards.append(m.convert_to_value("黑桃", 2))
# print(cbn.has_straight(tmp_cards))

# # has_gold_three_of_a_kind
# tmp_cards = []
# tmp_cards.append(m.convert_to_value("黑桃", 10))
# tmp_cards.append(m.convert_to_value("黑桃", 11))
# tmp_cards.append(m.convert_to_value("愛心", 7))
# tmp_cards.append(m.convert_to_value("梅花", 7))
# tmp_cards.append(m.convert_to_value("方塊", 7))
# tmp_cards.append(m.convert_to_value("方塊", 6))
# tmp_cards.append(m.convert_to_value("方塊", 9))
# tmp_cards.append(m.convert_to_value("方塊", 4))
# tmp_cards.append(m.convert_to_value("愛心", 10))
# tmp_cards.append(m.convert_to_value("黑桃", 2))
# gold = 7
# print(cbn.has_gold_three_of_a_kind(tmp_cards, gold))

# # has_three_of_a_kind
# tmp_cards = []
# tmp_cards.append(m.convert_to_value("黑桃", 1))
# tmp_cards.append(m.convert_to_value("愛心", 4))
# tmp_cards.append(m.convert_to_value("愛心", 3))
# tmp_cards.append(m.convert_to_value("梅花", 3))
# tmp_cards.append(m.convert_to_value("方塊", 3))
# tmp_cards.append(m.convert_to_value("方塊", 6))
# tmp_cards.append(m.convert_to_value("方塊", 9))
# tmp_cards.append(m.convert_to_value("方塊", 4))
# tmp_cards.append(m.convert_to_value("愛心", 10))
# tmp_cards.append(m.convert_to_value("黑桃", 2))
# print(cbn.has_three_of_a_kind(tmp_cards))

# # has_two_pairs
# tmp_cards = []
# tmp_cards.append(m.convert_to_value("黑桃", 1))
# tmp_cards.append(m.convert_to_value("愛心", 1))
# tmp_cards.append(m.convert_to_value("愛心", 4))
# tmp_cards.append(m.convert_to_value("梅花", 3))
# tmp_cards.append(m.convert_to_value("方塊", 3))
# tmp_cards.append(m.convert_to_value("方塊", 6))
# tmp_cards.append(m.convert_to_value("方塊", 9))
# tmp_cards.append(m.convert_to_value("方塊", 4))
# tmp_cards.append(m.convert_to_value("愛心", 10))
# tmp_cards.append(m.convert_to_value("黑桃", 2))
# print(cbn.has_two_pairs(tmp_cards))

# # has_one_pair
# tmp_cards = []
# tmp_cards.append(m.convert_to_value("黑桃", 1))
# tmp_cards.append(m.convert_to_value("愛心", 2))
# tmp_cards.append(m.convert_to_value("愛心", 4))
# tmp_cards.append(m.convert_to_value("梅花", 3))
# tmp_cards.append(m.convert_to_value("方塊", 3))
# tmp_cards.append(m.convert_to_value("方塊", 6))
# tmp_cards.append(m.convert_to_value("方塊", 9))
# tmp_cards.append(m.convert_to_value("方塊", 4))
# tmp_cards.append(m.convert_to_value("愛心", 10))
# tmp_cards.append(m.convert_to_value("黑桃", 2))
# print(cbn.has_one_pair(tmp_cards))

# # has_high_card
# tmp_cards = []
# tmp_cards.append(m.convert_to_value("黑桃", 1))
# tmp_cards.append(m.convert_to_value("愛心", 2))
# tmp_cards.append(m.convert_to_value("愛心", 3))
# tmp_cards.append(m.convert_to_value("梅花", 6))
# tmp_cards.append(m.convert_to_value("方塊", 7))
# print(cbn.has_high_card(tmp_cards))
#
# # has_high_card
# tmp_cards = []
# print(cbn.has_high_card(tmp_cards))
