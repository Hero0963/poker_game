import cp_const as c
import cp_check_poker_hand as m
import cp_poker_hand_separate_cards as sc
from typing import List
import collections
import copy
import itertools

# get_five_of_a_kind_from_poker_hand
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
# tmp_cards.append(m.convert_to_value("黑桃", 4))
# tmp_cards.append(m.convert_to_value("黑桃", 4))
# tmp_cards.append(m.convert_to_value("愛心", 4))
# tmp_cards.append(m.convert_to_value("梅花", 4))
# tmp_cards.append(m.convert_to_value("方塊", 4))
# print(sc.get_five_of_a_kind_from_poker_hand(tmp_cards))

# # # get_straight_flush_from_poker_hand
# tmp_cards = []
# tmp_cards.append(m.convert_to_value("黑桃", 7))
# tmp_cards.append(m.convert_to_value("黑桃", 8))
# tmp_cards.append(m.convert_to_value("黑桃", 9))
# tmp_cards.append(m.convert_to_value("黑桃", 10))
# tmp_cards.append(m.convert_to_value("黑桃", 11))
# tmp_cards.append(m.convert_to_value("方塊", 7))
# tmp_cards.append(m.convert_to_value("方塊", 4))
# tmp_cards.append(m.convert_to_value("方塊", 6))
# tmp_cards.append(m.convert_to_value("方塊", 5))
# tmp_cards.append(m.convert_to_value("方塊", 8))
# tmp_cards.append(m.convert_to_value("方塊", 9))
# tmp_cards.append(m.convert_to_value("方塊", 10))
# tmp_cards.append(m.convert_to_value("方塊", 11))
# print(sc.get_straight_flush_from_poker_hand(tmp_cards))


# # # get_five_of_gold_tiger_from_poker_hand
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
# print(sc.get_five_of_gold_tiger_from_poker_hand(tmp_cards, gold))


# # has_four_of_a_kind
# tmp_cards = []
# tmp_cards.append(m.convert_to_value("黑桃", 7))
# tmp_cards.append(m.convert_to_value("愛心", 7))
# tmp_cards.append(m.convert_to_value("方塊", 7))
# tmp_cards.append(m.convert_to_value("梅花", 7))
# tmp_cards.append(m.convert_to_value("愛心", 5))
# tmp_cards.append(m.convert_to_value("黑桃", 5))
# tmp_cards.append(m.convert_to_value("黑桃", 6))
# tmp_cards.append(m.convert_to_value("方塊", 5))
# tmp_cards.append(m.convert_to_value("梅花", 3))
# tmp_cards.append(m.convert_to_value("愛心", 4))
# target = sc.get_four_of_a_kind_from_poker_hand(tmp_cards)
# print(target)
# print(len(target))


# # # get_gold_full_house_from_poker_hand
# tmp_cards = []
# tmp_cards.append(m.convert_to_value("黑桃", 7))
# tmp_cards.append(m.convert_to_value("方塊", 7))
# tmp_cards.append(m.convert_to_value("愛心", 7))
# tmp_cards.append(m.convert_to_value("梅花", 7))
# tmp_cards.append(m.convert_to_value("梅花", 5))
# tmp_cards.append(m.convert_to_value("方塊", 5))
# tmp_cards.append(m.convert_to_value("黑桃", 5))
# tmp_cards.append(m.convert_to_value("愛心", 5))
# tmp_cards.append(m.convert_to_value("愛心", 10))
# tmp_cards.append(m.convert_to_value("梅花", 10))
# tmp_cards.append(m.convert_to_value("方塊", 10))
# gold = 7
# target = sc.get_gold_full_house_from_poker_hand(tmp_cards, gold)
# print(target)
# print(len(target))


# # get_full_house_from_poker_hand
# tmp_cards = []
# tmp_cards.append(m.convert_to_value("黑桃", 7))
# tmp_cards.append(m.convert_to_value("方塊", 7))
# tmp_cards.append(m.convert_to_value("愛心", 7))
# tmp_cards.append(m.convert_to_value("梅花", 5))
# tmp_cards.append(m.convert_to_value("方塊", 5))
# tmp_cards.append(m.convert_to_value("黑桃", 5))
# tmp_cards.append(m.convert_to_value("愛心", 5))
# tmp_cards.append(m.convert_to_value("愛心", 10))
# tmp_cards.append(m.convert_to_value("梅花", 10))
# tmp_cards.append(m.convert_to_value("方塊", 10))
# gold = 6
# target = sc.get_full_house_from_poker_hand(tmp_cards)
# print(target)
# print(len(target))
#
# """
# 777 55 = 6
# 777 T T = 3
# 555 77 = 4 * 3 = 12
# 555 TT = 4 * 3 = 12
# TTT 77 = 3
# TTT 55 = 6
#
# """


# # get_pair_with_flush_from_poker_hand
# tmp_cards = []
# tmp_cards.append(m.convert_to_value("黑桃", 7))
# tmp_cards.append(m.convert_to_value("黑桃", 7))
# tmp_cards.append(m.convert_to_value("黑桃", 2))
# tmp_cards.append(m.convert_to_value("黑桃", 6))
# tmp_cards.append(m.convert_to_value("黑桃", 5))
# tmp_cards.append(m.convert_to_value("方塊", 6))
# tmp_cards.append(m.convert_to_value("黑桃", 11))
# tmp_cards.append(m.convert_to_value("方塊", 2))
# tmp_cards.append(m.convert_to_value("黑桃", 3))
# tmp_cards.append(m.convert_to_value("方塊", 5))
# tmp_cards.append(m.convert_to_value("方塊", 7))
# tmp_cards.append(m.convert_to_value("黑桃", 8))
# target = sc.get_pair_with_flush_from_poker_hand(tmp_cards)
# print(target)
# print(len(target))


# # # get_flush_from_poker_hand
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
# tmp_cards.append(m.convert_to_value("方塊", 6))
# tmp_cards.append(m.convert_to_value("方塊", 7))
# tmp_cards.append(m.convert_to_value("方塊", 3))
# target = sc.get_flush_from_poker_hand(tmp_cards)
# print(target)
# print(len(target))


# # get_straight_from_poker_hand
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
# target = sc.get_straight_from_poker_hand(tmp_cards)
# print(target)
# print(len(target))
# # """
# # 9 T J Q K = 2
# # T J Q K A = 2
# # """
#
#
# # get_gold_three_of_a_kind_from_poker_hand
# tmp_cards = []
# tmp_cards.append(m.convert_to_value("黑桃", 10))
# tmp_cards.append(m.convert_to_value("黑桃", 11))
# tmp_cards.append(m.convert_to_value("黑桃", 7))
# tmp_cards.append(m.convert_to_value("愛心", 7))
# tmp_cards.append(m.convert_to_value("方塊", 7))
# tmp_cards.append(m.convert_to_value("方塊", 6))
# tmp_cards.append(m.convert_to_value("方塊", 9))
# tmp_cards.append(m.convert_to_value("方塊", 4))
# tmp_cards.append(m.convert_to_value("愛心", 10))
# tmp_cards.append(m.convert_to_value("黑桃", 2))
# gold = 7
# target = sc.get_gold_three_of_a_kind_from_poker_hand(tmp_cards, gold)
# print(target)
# print(len(target))


# get_three_of_a_kind_from_poker_hand
# tmp_cards = []
# tmp_cards.append(m.convert_to_value("黑桃", 1))
# tmp_cards.append(m.convert_to_value("愛心", 4))
# tmp_cards.append(m.convert_to_value("愛心", 3))
# tmp_cards.append(m.convert_to_value("梅花", 3))
# tmp_cards.append(m.convert_to_value("方塊", 3))
# tmp_cards.append(m.convert_to_value("黑桃", 3))
# tmp_cards.append(m.convert_to_value("方塊", 6))
# tmp_cards.append(m.convert_to_value("方塊", 9))
# tmp_cards.append(m.convert_to_value("方塊", 4))
# tmp_cards.append(m.convert_to_value("愛心", 10))
# tmp_cards.append(m.convert_to_value("黑桃", 2))
# tmp_cards.append(m.convert_to_value("梅花", 4))
# target = sc.get_three_of_a_kind_from_poker_hand(tmp_cards)
# print(target)
# print(len(target))


# # # get_two_pairs_from_poker_hand
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
# tmp_cards.append(m.convert_to_value("黑桃", 4))
# target = sc.get_two_pairs_from_poker_hand(tmp_cards)
# print(target)
# print(len(target))


# # # get_one_pair_from_poker_hand
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
# target = sc.get_one_pair_from_poker_hand(tmp_cards)
# print(target)
# print(len(target))
#
#
# # get_high_card_from_poker_hand
# tmp_cards = []
# tmp_cards.append(m.convert_to_value("黑桃", 1))
# tmp_cards.append(m.convert_to_value("愛心", 2))
# tmp_cards.append(m.convert_to_value("愛心", 3))
# tmp_cards.append(m.convert_to_value("梅花", 6))
# tmp_cards.append(m.convert_to_value("方塊", 7))
# target = sc.get_high_card_from_poker_hand(tmp_cards)
# print(target)
# print(len(target))
