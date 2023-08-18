import cp_const as c
import cp_check_poker_hand as m
import cp_calculate_combat as cal

# new_offset = cal.update_offset(12, 5)
# print(new_offset)
#
#
# # get_score_five_of_a_kind
# tmp_cards = []
# tmp_cards.append(m.convert_to_value("黑桃", 7))
# tmp_cards.append(m.convert_to_value("黑桃", 7))
# tmp_cards.append(m.convert_to_value("愛心", 7))
# tmp_cards.append(m.convert_to_value("梅花", 7))
# tmp_cards.append(m.convert_to_value("方塊", 7))
# print(m.is_five_of_a_kind(tmp_cards))
# print(cal.get_score_five_of_a_kind(tmp_cards))
#
#
# # get_score_straight_flush
# tmp_cards = []
# tmp_cards.append(m.convert_to_value("黑桃", 7))
# tmp_cards.append(m.convert_to_value("黑桃", 8))
# tmp_cards.append(m.convert_to_value("黑桃", 9))
# tmp_cards.append(m.convert_to_value("黑桃", 10))
# tmp_cards.append(m.convert_to_value("黑桃", 11))
# print(m.is_straight_flush(tmp_cards))
# print(cal.get_score_straight_flush(tmp_cards))
#
#
# # get_score_five_gold_tiger
# tmp_cards = []
# tmp_cards.append(m.convert_to_value("黑桃", 6))
# tmp_cards.append(m.convert_to_value("黑桃", 6))
# tmp_cards.append(m.convert_to_value("梅花", 6))
# tmp_cards.append(m.convert_to_value("方塊", 6))
# tmp_cards.append(m.convert_to_value("黑桃", 11))
# gold = 6
# print(m.is_five_of_gold_tiger(tmp_cards, gold))
# print(cal.get_score_five_gold_tiger(tmp_cards, gold))
#
#
# # # get_score_four_of_a_kind
# tmp_cards = []
# tmp_cards.append(m.convert_to_value("黑桃", 7))
# tmp_cards.append(m.convert_to_value("愛心", 7))
# tmp_cards.append(m.convert_to_value("方塊", 7))
# tmp_cards.append(m.convert_to_value("梅花", 7))
# tmp_cards.append(m.convert_to_value("愛心", 5))
# print(m.is_four_of_a_kind(tmp_cards))
# print(cal.get_score_four_of_a_kind(tmp_cards))
#
#
# # get_score_gold_full_house
# tmp_cards = []
# tmp_cards.append(m.convert_to_value("黑桃", 7))
# tmp_cards.append(m.convert_to_value("方塊", 7))
# tmp_cards.append(m.convert_to_value("愛心", 7))
# tmp_cards.append(m.convert_to_value("梅花", 5))
# tmp_cards.append(m.convert_to_value("方塊", 5))
# gold = 7
# print(m.is_gold_full_house(tmp_cards, gold))
# print(cal.get_score_gold_full_house(tmp_cards, gold))

# # get_score_full_house
# tmp_cards = []
# tmp_cards.append(m.convert_to_value("黑桃", 7))
# tmp_cards.append(m.convert_to_value("方塊", 7))
# tmp_cards.append(m.convert_to_value("愛心", 7))
# tmp_cards.append(m.convert_to_value("梅花", 5))
# tmp_cards.append(m.convert_to_value("方塊", 5))
# gold = 6
# print(m.is_full_house(tmp_cards))
# print(cal.get_score_full_house(tmp_cards))

#
# # get_score_pair_with_flush
# tmp_cards = []
# tmp_cards.append(m.convert_to_value("黑桃", 7))
# tmp_cards.append(m.convert_to_value("黑桃", 7))
# tmp_cards.append(m.convert_to_value("黑桃", 2))
# tmp_cards.append(m.convert_to_value("黑桃", 6))
# tmp_cards.append(m.convert_to_value("黑桃", 5))
# print(m.is_pair_with_flush(tmp_cards))
# print(cal.get_score_pair_with_flush(tmp_cards))
#
# # # get_score_flush
# tmp_cards = []
# tmp_cards.append(m.convert_to_value("黑桃", 9))
# tmp_cards.append(m.convert_to_value("黑桃", 6))
# tmp_cards.append(m.convert_to_value("黑桃", 5))
# tmp_cards.append(m.convert_to_value("黑桃", 4))
# tmp_cards.append(m.convert_to_value("黑桃", 3))
# print(m.is_flush(tmp_cards))
# print(cal.get_score_flush(tmp_cards))
#
#
# ## get_score_straight
# tmp_cards = []
# tmp_cards.append(m.convert_to_value("黑桃", 10))
# tmp_cards.append(m.convert_to_value("黑桃", 11))
# tmp_cards.append(m.convert_to_value("愛心", 12))
# tmp_cards.append(m.convert_to_value("梅花", 13))
# tmp_cards.append(m.convert_to_value("方塊", 1))
# print(m.is_straight(tmp_cards))
# print(cal.get_score_straight(tmp_cards))
#
#
# # # get_score_gold_three_of_a_kind
# tmp_cards = []
# tmp_cards.append(m.convert_to_value("黑桃", 10))
# tmp_cards.append(m.convert_to_value("黑桃", 11))
# tmp_cards.append(m.convert_to_value("愛心", 7))
# tmp_cards.append(m.convert_to_value("梅花", 7))
# tmp_cards.append(m.convert_to_value("方塊", 7))
# gold = 7
# print(m.is_gold_three_of_a_kind(tmp_cards, gold))
# print(cal.get_score_gold_three_of_a_kind(tmp_cards, gold))
#
#
# # # get_score_three_of_a_kind
# tmp_cards = []
# tmp_cards.append(m.convert_to_value("黑桃", 1))
# tmp_cards.append(m.convert_to_value("愛心", 4))
# tmp_cards.append(m.convert_to_value("愛心", 3))
# tmp_cards.append(m.convert_to_value("梅花", 3))
# tmp_cards.append(m.convert_to_value("方塊", 3))
# print(m.is_three_of_a_kind(tmp_cards))
# print(cal.get_score_three_of_a_kind(tmp_cards))
#
#
# # get_score_two_pairs
# tmp_cards = []
# tmp_cards.append(m.convert_to_value("黑桃", 1))
# tmp_cards.append(m.convert_to_value("愛心", 1))
# tmp_cards.append(m.convert_to_value("愛心", 4))
# tmp_cards.append(m.convert_to_value("梅花", 3))
# tmp_cards.append(m.convert_to_value("方塊", 3))
# print(m.is_two_pairs(tmp_cards))
# print(cal.get_score_two_pairs(tmp_cards))
#
#
# # get_score_one_pair
# tmp_cards = []
# tmp_cards.append(m.convert_to_value("黑桃", 1))
# tmp_cards.append(m.convert_to_value("愛心", 2))
# tmp_cards.append(m.convert_to_value("愛心", 4))
# tmp_cards.append(m.convert_to_value("梅花", 3))
# tmp_cards.append(m.convert_to_value("方塊", 3))
# print(m.is_one_pair(tmp_cards))
# print(cal.get_score_one_pair(tmp_cards))
#
#
# # get_score_high_card
# tmp_cards = []
# tmp_cards.append(m.convert_to_value("黑桃", 1))
# tmp_cards.append(m.convert_to_value("愛心", 2))
# tmp_cards.append(m.convert_to_value("愛心", 3))
# tmp_cards.append(m.convert_to_value("梅花", 6))
# tmp_cards.append(m.convert_to_value("方塊", 7))
# print(m.is_high_card(tmp_cards))
# print(cal.get_score_high_card(tmp_cards))
