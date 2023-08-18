import collections

import cp_const as c
import cp_check_poker_hand as m
from typing import List

PH = c.PokerHand
COMBAT_BASE = c.COMBAT_BASE


def update_offset(ori_offset: int, x: int) -> int:
    offset_base = 100
    return ori_offset * offset_base + x


def get_score_five_of_a_kind(cards: List[int]) -> int:
    base_score = COMBAT_BASE * PH.FIVE_HIGH
    offset = 0
    card = cards[0]
    rank = m.get_card_rank(card)
    rk_value = m.convert_rank_value(rank)

    offset = update_offset(offset, rk_value)
    return base_score + offset


def get_score_straight_flush(cards: List[int]) -> int:
    base_score = COMBAT_BASE * PH.STRAIGHT_FLUSH
    offset = 0
    record_rk_values = []
    for card in cards:
        rank = m.get_card_rank(card)
        rk_value = m.convert_rank_value(rank)
        record_rk_values.append(rk_value)

    record_rk_values.sort()
    max_rk_value = record_rk_values[-1]
    offset = update_offset(offset, max_rk_value)
    return base_score + offset


def get_score_five_gold_tiger(cards: List[int], gold: int) -> int:
    base_score = COMBAT_BASE * PH.FIVE_OF_GOLD_TIGER
    offset = 0
    offset = update_offset(offset, gold)
    single_value = 0
    for card in cards:
        rank = m.get_card_rank(card)
        rk_value = m.convert_rank_value(rank)
        if rk_value != gold:
            single_value = rk_value
            break

    offset = update_offset(offset, single_value)
    return base_score + offset


def get_score_four_of_a_kind(cards: List[int]) -> int:
    base_score = COMBAT_BASE * PH.FOUR_OF_A_KIND
    offset = 0
    record_rk_value = collections.defaultdict(int)
    for card in cards:
        rank = m.get_card_rank(card)
        rk_value = m.convert_rank_value(rank)
        record_rk_value[rk_value] += 1
    kind_value, single_value = 0, 0
    for rk_value, count in record_rk_value.items():
        if count >= 4:
            kind_value = rk_value
        else:
            single_value = rk_value

    if single_value == 0:
        single_value = kind_value
    offset = update_offset(offset, kind_value)
    offset = update_offset(offset, single_value)
    return base_score + offset


def get_score_gold_full_house(cards: List[int], gold: int) -> int:
    base_score = COMBAT_BASE * PH.GOLD_FULL_HOUSE
    offset = 0
    offset = update_offset(offset, gold)

    record_rk_value = collections.defaultdict(int)
    for card in cards:
        rank = m.get_card_rank(card)
        rk_value = m.convert_rank_value(rank)
        record_rk_value[rk_value] += 1

    for rk_value, count in record_rk_value.items():
        if rk_value != gold and count >= 2:
            offset = update_offset(offset, rk_value)

    return base_score + offset


def get_score_full_house(cards: List[int]) -> int:
    base_score = COMBAT_BASE * PH.FULL_HOUSE
    offset = 0

    record_rk_value = collections.defaultdict(int)
    for card in cards:
        rank = m.get_card_rank(card)
        rk_value = m.convert_rank_value(rank)
        record_rk_value[rk_value] += 1

    update_order = [0, 0]
    for rk_value, count in record_rk_value.items():
        if count == 3:
            update_order[0] = rk_value
        else:
            update_order[1] = rk_value

    for v in update_order:
        offset = update_offset(offset, v)

    return base_score + offset


def get_score_pair_with_flush(cards: List[int]) -> int:
    base_score = COMBAT_BASE * PH.PAIR_WITH_FLUSH
    offset = 0

    record_rk_values = collections.defaultdict(int)
    pair_value = -1
    for card in cards:
        rank = m.get_card_rank(card)
        rk_value = m.convert_rank_value(rank)
        record_rk_values[rk_value] += 1
        if record_rk_values[rk_value] > 1:
            pair_value = max(pair_value, rk_value)

    offset = update_offset(offset, pair_value)
    values = []
    for rk_value, count in record_rk_values.items():
        if rk_value != pair_value:
            values.append(rk_value)

    values.sort(reverse=True)
    for v in values:
        offset = update_offset(offset, v)

    return base_score + offset


def get_score_flush(cards: List[int]) -> int:
    base_score = COMBAT_BASE * PH.FLUSH
    offset = 0
    rk_values = []
    for card in cards:
        rank = m.get_card_rank(card)
        rk_value = m.convert_rank_value(rank)
        rk_values.append(rk_value)

    rk_values.sort(reverse=True)
    for x in rk_values:
        offset = update_offset(offset, x)

    return base_score + offset


def get_score_straight(cards: List[int]) -> int:
    base_score = COMBAT_BASE * PH.STRAIGHT
    offset = 0
    rk_values = []
    for card in cards:
        rank = m.get_card_rank(card)
        rk_value = m.convert_rank_value(rank)
        rk_values.append(rk_value)

    rk_values.sort(reverse=True)
    for x in rk_values:
        offset = update_offset(offset, x)
    return base_score + offset


def get_score_gold_three_of_a_kind(cards: List[int], gold: int) -> int:
    base_score = COMBAT_BASE * PH.GOLD_THREE_OF_A_KIND
    offset = 0
    offset = update_offset(offset, gold)

    rk_values = []
    for card in cards:
        rank = m.get_card_rank(card)
        rk_value = m.convert_rank_value(rank)
        if rk_value != gold:
            rk_values.append(rk_value)

    rk_values.sort(reverse=True)
    for x in rk_values:
        offset = update_offset(offset, x)
    return base_score + offset


def get_score_three_of_a_kind(cards: List[int]) -> int:
    base_score = COMBAT_BASE * PH.THREE_OF_A_KIND
    offset = 0

    record_rk_value = collections.defaultdict(int)
    for card in cards:
        rank = m.get_card_rank(card)
        rk_value = m.convert_rank_value(rank)
        record_rk_value[rk_value] += 1

    kind_value = 0
    rk_values = []
    for rk_value, count in record_rk_value.items():
        if count >= 3:
            kind_value = rk_value
        else:
            rk_values.append(rk_value)

    while len(rk_values) < 2:
        rk_values.append(kind_value)

    offset = update_offset(offset, kind_value)
    rk_values.sort(reverse=True)
    for x in rk_values:
        offset = update_offset(offset, x)
    return base_score + offset


def get_score_two_pairs(cards: List[int]) -> int:
    base_score = COMBAT_BASE * PH.TWO_PAIRS
    offset = 0

    record_rk_value = collections.defaultdict(int)
    for card in cards:
        rank = m.get_card_rank(card)
        rk_value = m.convert_rank_value(rank)
        record_rk_value[rk_value] += 1

    pair_values = []
    single_value = 0
    for rk_value, count in record_rk_value.items():
        if count >= 2:
            pair_values.append(rk_value)
        else:
            single_value = rk_value

    pair_values.sort(reverse=True)
    for x in pair_values:
        offset = update_offset(offset, x)

    offset = update_offset(offset, single_value)
    return base_score + offset


def get_score_one_pair(cards: List[int]) -> int:
    base_score = COMBAT_BASE * PH.ONE_PAIR
    offset = 0

    record_rk_value = collections.defaultdict(int)
    for card in cards:
        rank = m.get_card_rank(card)
        rk_value = m.convert_rank_value(rank)
        record_rk_value[rk_value] += 1

    pair_value = 0
    rk_values = []
    for rk_value, count in record_rk_value.items():
        if count >= 2:
            pair_value = rk_value
        else:
            rk_values.append(rk_value)

    offset = update_offset(offset, pair_value)
    rk_values.sort(reverse=True)
    for x in rk_values:
        offset = update_offset(offset, x)
    return base_score + offset


def get_score_high_card(cards: List[int]) -> int:
    base_score = COMBAT_BASE * PH.HIGH_CARD
    offset = 0

    rk_values = []
    for card in cards:
        rank = m.get_card_rank(card)
        rk_value = m.convert_rank_value(rank)
        rk_values.append(rk_value)

    rk_values.sort(reverse=True)
    for x in rk_values:
        offset = update_offset(offset, x)

    return base_score + offset
