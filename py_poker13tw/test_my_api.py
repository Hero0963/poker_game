import cp_const as c
import cp_check_poker_hand as m
import cp_calculate_combat as cal
import cp_poker_hand_combinations as cbn
from typing import List, Tuple
import collections
import cp_api

# get_poker_hand_rank_and_score

# five_of_a_kind
# tmp_cards = []
# tmp_cards.append(m.convert_to_value("黑桃", 7))
# tmp_cards.append(m.convert_to_value("黑桃", 7))
# tmp_cards.append(m.convert_to_value("愛心", 7))
# tmp_cards.append(m.convert_to_value("梅花", 7))
# tmp_cards.append(m.convert_to_value("方塊", 7))
# gold = 5
# print(cp_api.get_poker_hand_rank_and_score(tmp_cards, gold))


# # straight_flush
# tmp_cards = []
# tmp_cards.append(m.convert_to_value("黑桃", 7))
# tmp_cards.append(m.convert_to_value("黑桃", 8))
# tmp_cards.append(m.convert_to_value("黑桃", 9))
# tmp_cards.append(m.convert_to_value("黑桃", 10))
# tmp_cards.append(m.convert_to_value("黑桃", 11))
# gold = 5
# print(cp_api.get_poker_hand_rank_and_score(tmp_cards, gold))

# # five_gold_tiger
# tmp_cards = []
# tmp_cards.append(m.convert_to_value("黑桃", 6))
# tmp_cards.append(m.convert_to_value("黑桃", 6))
# tmp_cards.append(m.convert_to_value("梅花", 6))
# tmp_cards.append(m.convert_to_value("方塊", 6))
# tmp_cards.append(m.convert_to_value("黑桃", 11))
# gold = 6
# print(cp_api.get_poker_hand_rank_and_score(tmp_cards, gold))

# four_of_a_kind
# tmp_cards = []
# tmp_cards.append(m.convert_to_value("黑桃", 7))
# tmp_cards.append(m.convert_to_value("愛心", 7))
# tmp_cards.append(m.convert_to_value("方塊", 7))
# tmp_cards.append(m.convert_to_value("梅花", 7))
# tmp_cards.append(m.convert_to_value("愛心", 5))
# gold = 6
# print(cp_api.get_poker_hand_rank_and_score(tmp_cards, gold))

# gold_full_house
# tmp_cards = []
# tmp_cards.append(m.convert_to_value("黑桃", 7))
# tmp_cards.append(m.convert_to_value("方塊", 7))
# tmp_cards.append(m.convert_to_value("愛心", 7))
# tmp_cards.append(m.convert_to_value("梅花", 5))
# tmp_cards.append(m.convert_to_value("方塊", 5))
# gold = 7
# print(cp_api.get_poker_hand_rank_and_score(tmp_cards, gold))

# # full_house
# tmp_cards = []
# tmp_cards.append(m.convert_to_value("黑桃", 7))
# tmp_cards.append(m.convert_to_value("方塊", 7))
# tmp_cards.append(m.convert_to_value("愛心", 7))
# tmp_cards.append(m.convert_to_value("梅花", 5))
# tmp_cards.append(m.convert_to_value("方塊", 5))
# gold = 6
# print(cp_api.get_poker_hand_rank_and_score(tmp_cards, gold))


# # pair_with_flush
# tmp_cards = []
# tmp_cards.append(m.convert_to_value("黑桃", 7))
# tmp_cards.append(m.convert_to_value("黑桃", 7))
# tmp_cards.append(m.convert_to_value("黑桃", 5))
# tmp_cards.append(m.convert_to_value("黑桃", 5))
# tmp_cards.append(m.convert_to_value("黑桃", 2))
# gold = 8
# print(cp_api.get_poker_hand_rank_and_score(tmp_cards, gold))


# # flush
# tmp_cards = []
# tmp_cards.append(m.convert_to_value("黑桃", 9))
# tmp_cards.append(m.convert_to_value("黑桃", 6))
# tmp_cards.append(m.convert_to_value("黑桃", 5))
# tmp_cards.append(m.convert_to_value("黑桃", 4))
# tmp_cards.append(m.convert_to_value("黑桃", 3))
# gold = 8
# print(cp_api.get_poker_hand_rank_and_score(tmp_cards, gold))


# straight
# tmp_cards = []
# tmp_cards.append(m.convert_to_value("黑桃", 10))
# tmp_cards.append(m.convert_to_value("黑桃", 11))
# tmp_cards.append(m.convert_to_value("愛心", 12))
# tmp_cards.append(m.convert_to_value("梅花", 13))
# tmp_cards.append(m.convert_to_value("方塊", 1))
# gold = 8
# print(cp_api.get_poker_hand_rank_and_score(tmp_cards, gold))


# # gold_three_of_a_kind
# tmp_cards = []
# tmp_cards.append(m.convert_to_value("黑桃", 10))
# tmp_cards.append(m.convert_to_value("黑桃", 11))
# tmp_cards.append(m.convert_to_value("愛心", 7))
# tmp_cards.append(m.convert_to_value("梅花", 7))
# tmp_cards.append(m.convert_to_value("方塊", 7))
# gold = 7
# print(cp_api.get_poker_hand_rank_and_score(tmp_cards, gold))


# three_of_a_kind
# tmp_cards = []
# tmp_cards.append(m.convert_to_value("黑桃", 1))
# tmp_cards.append(m.convert_to_value("愛心", 4))
# tmp_cards.append(m.convert_to_value("愛心", 3))
# tmp_cards.append(m.convert_to_value("梅花", 3))
# tmp_cards.append(m.convert_to_value("方塊", 3))
# gold = 7
# print(cp_api.get_poker_hand_rank_and_score(tmp_cards, gold))



# # two_pairs
# tmp_cards = []
# tmp_cards.append(m.convert_to_value("黑桃", 1))
# tmp_cards.append(m.convert_to_value("愛心", 1))
# tmp_cards.append(m.convert_to_value("愛心", 4))
# tmp_cards.append(m.convert_to_value("梅花", 3))
# tmp_cards.append(m.convert_to_value("方塊", 3))
# gold = 7
# print(cp_api.get_poker_hand_rank_and_score(tmp_cards, gold))

# one_pair
# tmp_cards = []
# tmp_cards.append(m.convert_to_value("黑桃", 1))
# tmp_cards.append(m.convert_to_value("愛心", 2))
# tmp_cards.append(m.convert_to_value("愛心", 4))
# tmp_cards.append(m.convert_to_value("梅花", 3))
# tmp_cards.append(m.convert_to_value("方塊", 3))
# gold = 7
# print(cp_api.get_poker_hand_rank_and_score(tmp_cards, gold))


# # high_card
# tmp_cards = []
# tmp_cards.append(m.convert_to_value("黑桃", 1))
# tmp_cards.append(m.convert_to_value("愛心", 2))
# tmp_cards.append(m.convert_to_value("愛心", 3))
# tmp_cards.append(m.convert_to_value("梅花", 6))
# tmp_cards.append(m.convert_to_value("方塊", 7))
# gold = 7
# print(cp_api.get_poker_hand_rank_and_score(tmp_cards, gold))