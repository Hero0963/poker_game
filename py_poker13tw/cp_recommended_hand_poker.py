import cp_const as c
import cp_check_poker_hand as m
import cp_poker_hand_separate_cards as csc
from typing import List
import collections
import cp_api
import copy
import itertools

"""
單一指定墩贏某家且滿足牌型，輸的玩家需支付以下分數

頭墩
衝三 + 1
彩金衝三 +2


中墩
彩金三條 +1
彩金葫蘆 +2
鐵支 +2
同花順 +2
彩金五虎將 +5
五枚 +10


尾墩
彩金三條 +1
彩金葫蘆 +1
鐵支 +1
同花順 +1
彩金五虎將 +5
五枚 +5

"""


def get_recommend_poker(hc, tc, mc, gold):
    poker = [hc[:], tc[:], mc[:]]
    poker.sort(key=lambda pk: cp_api.get_poker_hand_rank_and_score(pk, gold)[-1])
    return poker


def top_up_cards(cur_poker_hand, unused_card, number):
    while len(cur_poker_hand) < number:
        cur_poker_hand.append(unused_card.pop())

    return cur_poker_hand


def get_sorted_unused_cards(cards):
    sorted_cards = cards[:]
    sorted_cards.sort(key=lambda card: m.convert_rank_value(m.get_card_rank(card)))
    return sorted_cards


def update_unused_cards(used_cards, original_cards):
    updated_unused_cards = []
    record = collections.Counter(original_cards)
    for uc in used_cards:
        record[uc] -= 1

    for card, count in record.items():
        if count > 0:
            for i in range(count):
                updated_unused_cards.append(card)

    return updated_unused_cards


def get_head_core_cards(cards: List[int], gold: int):
    gold_three_of_a_kind = csc.get_gold_three_of_a_kind_from_poker_hand(cards[:], gold)
    if gold_three_of_a_kind:
        return gold_three_of_a_kind

    three_of_a_kind = csc.get_three_of_a_kind_from_poker_hand(cards[:])
    if three_of_a_kind:
        return three_of_a_kind

    one_pair = csc.get_one_pair_from_poker_hand(cards[:])
    if one_pair:
        return one_pair

    high_card = csc.get_high_card_from_poker_hand(cards[:])
    if high_card:
        return high_card


def get_middle_tail_core_cards(cards: List[int], gold: int):
    five_of_a_kind = csc.get_five_of_a_kind_from_poker_hand(cards[:])
    if five_of_a_kind:
        return five_of_a_kind

    five_of_gold_tiger = csc.get_five_of_gold_tiger_from_poker_hand(cards[:], gold)
    if five_of_gold_tiger:
        return five_of_gold_tiger

    straight_flush = csc.get_straight_flush_from_poker_hand(cards[:])
    four_of_a_kind = csc.get_four_of_a_kind_from_poker_hand(cards[:])
    if straight_flush or four_of_a_kind:
        sf_fk = []
        if straight_flush:
            for poker in straight_flush:
                sf_fk.append(poker)
        if four_of_a_kind:
            for poker in four_of_a_kind:
                sf_fk.append(poker)
        return sf_fk

    gold_full_house = csc.get_gold_full_house_from_poker_hand(cards[:], gold)
    if gold_full_house:
        return gold_full_house

    gold_three_of_a_kind = csc.get_gold_three_of_a_kind_from_poker_hand(cards[:], gold)
    if gold_three_of_a_kind:
        return gold_three_of_a_kind

    full_house = csc.get_full_house_from_poker_hand(cards[:])
    if full_house:
        return full_house

    pair_with_flush = csc.get_pair_with_flush_from_poker_hand(cards[:])
    if pair_with_flush:
        return pair_with_flush

    flush = csc.get_flush_from_poker_hand(cards[:])
    if flush:
        return flush

    straight = csc.get_straight_from_poker_hand(cards[:])
    if straight:
        return straight

    three_of_a_kind = csc.get_three_of_a_kind_from_poker_hand(cards[:])
    if three_of_a_kind:
        return three_of_a_kind

    two_pairs = csc.get_two_pairs_from_poker_hand(cards[:])
    if two_pairs:
        return two_pairs

    one_pair = csc.get_one_pair_from_poker_hand(cards[:])
    if one_pair:
        return one_pair

    high_card = csc.get_high_card_from_poker_hand(cards[:])
    if high_card:
        return high_card


def order_553(cards: List[int], gold: int):
    result = []

    unused_cards = cards[:]
    middle_core = get_middle_tail_core_cards(unused_cards, gold)
    for mc in middle_core:
        unused_cards_m = update_unused_cards(used_cards=mc, original_cards=unused_cards)
        tail_core = get_middle_tail_core_cards(unused_cards_m, gold)
        for tc in tail_core:
            unused_cards_mt = update_unused_cards(used_cards=tc, original_cards=unused_cards_m)
            head_core = get_head_core_cards(unused_cards_mt, gold)
            for hc in head_core:
                unused_cards_mth = update_unused_cards(used_cards=hc, original_cards=unused_cards_mt)
                sorted_unused_cards = get_sorted_unused_cards(unused_cards_mth)
                completed_hc = top_up_cards(hc[:], sorted_unused_cards, 3)
                completed_mc = top_up_cards(mc[:], sorted_unused_cards, 5)
                completed_tc = top_up_cards(tc[:], sorted_unused_cards, 5)
                recommended_card = get_recommend_poker(completed_mc, completed_tc, completed_hc, gold)
                for rc in recommended_card:
                    ph_rank, combat = cp_api.get_poker_hand_rank_and_score(rc, gold)
                    # print("rc,ph_rank, combat =", rc, ph_rank, combat)

                if len(recommended_card[0]) == 3:
                    result.append(recommended_card)
    if len(result) > 20:
        result = result[0:20]
    return result


def order_535(cards: List[int], gold: int):
    result = []

    unused_cards = cards[:]
    middle_core = get_middle_tail_core_cards(unused_cards, gold)
    for mc in middle_core:
        unused_cards_m = update_unused_cards(used_cards=mc, original_cards=unused_cards)
        head_core = get_head_core_cards(unused_cards_m, gold)
        for hc in head_core:
            unused_cards_mh = update_unused_cards(used_cards=hc, original_cards=unused_cards_m)
            tail_core = get_middle_tail_core_cards(unused_cards_mh, gold)
            for tc in tail_core:
                unused_cards_mht = update_unused_cards(used_cards=tc, original_cards=unused_cards_mh)
                sorted_unused_cards = get_sorted_unused_cards(unused_cards_mht)
                top_up_cards(mc, sorted_unused_cards, 5)
                top_up_cards(tc, sorted_unused_cards, 5)
                top_up_cards(hc, sorted_unused_cards, 3)
                recommended_card = get_recommend_poker(hc, mc, tc, gold)
                for rc in recommended_card:
                    ph_rank, combat = cp_api.get_poker_hand_rank_and_score(rc, gold)
                    # print("rc,ph_rank, combat =", rc, ph_rank, combat)

                if len(recommended_card[0]) == 3:
                    result.append(recommended_card)

    if len(result) > 5:
        result = result[0:5]
    return result


def order_355(cards: List[int], gold: int):
    result = []

    unused_cards = cards[:]
    head_core = get_head_core_cards(unused_cards, gold)
    for hc in head_core:
        unused_cards_h = update_unused_cards(used_cards=hc, original_cards=unused_cards)
        middle_core = get_head_core_cards(unused_cards_h, gold)
        for mc in middle_core:
            unused_cards_hm = update_unused_cards(used_cards=mc, original_cards=unused_cards_h)
            tail_core = get_head_core_cards(unused_cards_hm, gold)
            for tc in tail_core:
                unused_cards_hmt = update_unused_cards(used_cards=tc, original_cards=unused_cards_hm)
                sorted_unused_cards = get_sorted_unused_cards(unused_cards_hmt)
                top_up_cards(mc, sorted_unused_cards, 5)
                top_up_cards(tc, sorted_unused_cards, 5)
                top_up_cards(hc, sorted_unused_cards, 3)
                recommended_card = get_recommend_poker(hc, mc, tc, gold)
                for rc in recommended_card:
                    ph_rank, combat = cp_api.get_poker_hand_rank_and_score(rc, gold)
                    # print("rc,ph_rank, combat =", rc, ph_rank, combat)

                if len(recommended_card[0]) == 3:
                    result.append(recommended_card)

    if len(result) > 5:
        result = result[0:5]
    return result
