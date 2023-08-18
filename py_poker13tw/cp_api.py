import cp_const as c
import cp_check_poker_hand as m
import cp_calculate_combat as cal
import cp_poker_hand_combinations as cbn
from typing import List, Tuple
import collections

PH = c.PokerHand()


def get_poker_hand_rank_and_score(cards: List[int], gold) -> Tuple[int, int]:
    if m.is_five_of_a_kind(cards):
        score = cal.get_score_five_of_a_kind(cards)
        return PH.FIVE_OF_GOLD_TIGER, score

    if m.is_straight_flush(cards):
        score = cal.get_score_straight_flush(cards)
        return PH.STRAIGHT_FLUSH, score

    if m.is_five_of_gold_tiger(cards, gold):
        score = cal.get_score_five_gold_tiger(cards, gold)
        return PH.FIVE_OF_GOLD_TIGER, score

    if m.is_four_of_a_kind(cards):
        score = cal.get_score_four_of_a_kind(cards)
        return PH.FOUR_OF_A_KIND, score

    if m.is_gold_full_house(cards, gold):
        score = cal.get_score_gold_full_house(cards, gold)
        return PH.GOLD_FULL_HOUSE, score

    if m.is_full_house(cards):
        score = cal.get_score_full_house(cards)
        return PH.FULL_HOUSE, score

    if m.is_pair_with_flush(cards):
        score = cal.get_score_pair_with_flush(cards)
        return PH.PAIR_WITH_FLUSH, score

    if m.is_flush(cards):
        score = cal.get_score_flush(cards)
        return PH.FLUSH, score

    if m.is_straight(cards):
        score = cal.get_score_straight(cards)
        return PH.STRAIGHT, score

    if m.is_gold_three_of_a_kind(cards, gold):
        score = cal.get_score_gold_three_of_a_kind(cards, gold)
        return PH.GOLD_THREE_OF_A_KIND, score

    if m.is_three_of_a_kind(cards):
        score = cal.get_score_three_of_a_kind(cards)
        return PH.THREE_OF_A_KIND, score

    if m.is_two_pairs(cards):
        score = cal.get_score_two_pairs(cards)
        return PH.TWO_PAIRS, score

    if m.is_one_pair(cards):
        score = cal.get_score_one_pair(cards)
        return PH.ONE_PAIR, score

    score = cal.get_score_high_card(cards)
    return PH.HIGH_CARD, score
