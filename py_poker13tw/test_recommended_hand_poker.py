import cp_recommended_hand_poker as rhp
import cp_const as c
import cp_check_poker_hand as m
import cp_poker_hand_separate_cards as csc
from typing import List
import collections
import copy
import itertools

# order_553

tmp_cards = []
tmp_cards.append(m.convert_to_value("黑桃", 2))
tmp_cards.append(m.convert_to_value("黑桃", 2))
tmp_cards.append(m.convert_to_value("方塊", 3))
tmp_cards.append(m.convert_to_value("黑桃", 4))
tmp_cards.append(m.convert_to_value("黑桃", 6))
tmp_cards.append(m.convert_to_value("愛心", 6))
tmp_cards.append(m.convert_to_value("方塊", 6))
tmp_cards.append(m.convert_to_value("梅花", 6))
tmp_cards.append(m.convert_to_value("黑桃", 7))
tmp_cards.append(m.convert_to_value("梅花", 8))
tmp_cards.append(m.convert_to_value("黑桃", 10))
tmp_cards.append(m.convert_to_value("愛心", 10))
tmp_cards.append(m.convert_to_value("黑桃", 12))
tmp_cards.append(m.convert_to_value("方塊", 12))
tmp_cards.append(m.convert_to_value("愛心", 1))
tmp_cards.append(m.convert_to_value("梅花", 1))
gold = 4
result = rhp.order_553(tmp_cards, gold)
print("test_553", result)
print(len(result))
for res in result:
    print("res =", res)
    d1, d2, d3 = res
    d = d1 + d2 + d3
    for card in d:
        suit = m.get_card_suit(card)
        rank = m.get_card_rank(card)
        print(suit, rank)
#
# # order_535
# tmp_cards = []
# tmp_cards.append(m.convert_to_value("黑桃", 1))
# tmp_cards.append(m.convert_to_value("愛心", 1))
# tmp_cards.append(m.convert_to_value("黑桃", 12))
# tmp_cards.append(m.convert_to_value("黑桃", 2))
# tmp_cards.append(m.convert_to_value("愛心", 2))
# tmp_cards.append(m.convert_to_value("梅花", 2))
# tmp_cards.append(m.convert_to_value("方塊", 11))
# tmp_cards.append(m.convert_to_value("梅花", 11))
# tmp_cards.append(m.convert_to_value("愛心", 3))
# tmp_cards.append(m.convert_to_value("方塊", 3))
# tmp_cards.append(m.convert_to_value("黑桃", 4))
# tmp_cards.append(m.convert_to_value("方塊", 4))
# tmp_cards.append(m.convert_to_value("梅花", 4))
# tmp_cards.append(m.convert_to_value("黑桃", 8))
# tmp_cards.append(m.convert_to_value("黑桃", 7))
# tmp_cards.append(m.convert_to_value("愛心", 6))
# gold = 4
# result = rhp.order_535(tmp_cards, gold)
# print("test_535", result)
# print(len(result))
#
# # order_355
# tmp_cards = []
# tmp_cards.append(m.convert_to_value("黑桃", 1))
# tmp_cards.append(m.convert_to_value("愛心", 1))
# tmp_cards.append(m.convert_to_value("黑桃", 12))
# tmp_cards.append(m.convert_to_value("黑桃", 2))
# tmp_cards.append(m.convert_to_value("愛心", 2))
# tmp_cards.append(m.convert_to_value("梅花", 2))
# tmp_cards.append(m.convert_to_value("方塊", 11))
# tmp_cards.append(m.convert_to_value("梅花", 11))
# tmp_cards.append(m.convert_to_value("愛心", 3))
# tmp_cards.append(m.convert_to_value("方塊", 3))
# tmp_cards.append(m.convert_to_value("黑桃", 4))
# tmp_cards.append(m.convert_to_value("方塊", 4))
# tmp_cards.append(m.convert_to_value("梅花", 4))
# tmp_cards.append(m.convert_to_value("黑桃", 8))
# tmp_cards.append(m.convert_to_value("黑桃", 7))
# tmp_cards.append(m.convert_to_value("愛心", 6))
# gold = 4
# result = rhp.order_355(tmp_cards, gold)
# print("test_355", result)
# print(len(result))
