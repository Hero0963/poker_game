import cp_const as c
import cp_check_poker_hand as m
from typing import List
import collections
import copy
import itertools

PH = c.PokerHand()


def has_five_of_a_kind(cards: List[int]) -> bool:
    if len(cards) < 5:
        return False

    record_rk_value = collections.defaultdict(int)
    for card in cards:
        rank = m.get_card_rank(card)
        rk_value = m.convert_rank_value(rank)
        record_rk_value[rk_value] += 1

    for rk_value, count in record_rk_value.items():
        if count >= 5:
            return True
    return False


def has_straight_flush(cards: List[int]) -> bool:
    if len(cards) < 5:
        return False

    def check_straight(poker: List[int]) -> bool:
        points = copy.deepcopy(poker)
        points.sort()

        cnt = 1
        pre_val = -1
        for p in points:
            if pre_val == -1:
                pre_val = p
                continue

            if p != pre_val + 1:
                cnt = 1
                pre_val = p
            else:
                cnt += 1
                pre_val = p

            if cnt >= 5:
                return True

        special_straight = [2, 3, 4, 5, 14]
        sp = set(points)
        for ss in special_straight:
            if ss not in sp:
                return False

        return True

    split_by_suit = collections.defaultdict(list)
    for card in cards:
        suit, rank = m.get_card_suit(card), m.get_card_rank(card)
        rk_value = m.convert_rank_value(rank)
        split_by_suit[suit].append(rk_value)

    for suit, values in split_by_suit.items():
        if check_straight(values):
            return True

    return False


def has_five_of_gold_tiger(cards: List[int], gold: int) -> bool:
    if len(cards) < 5:
        return False

    record_rk_value = collections.defaultdict(int)
    for card in cards:
        rank = m.get_card_rank(card)
        rk_value = m.convert_rank_value(rank)
        record_rk_value[rk_value] += 1

    if record_rk_value[gold] >= 4:
        return True

    return False


def has_four_of_a_kind(cards: List[int]) -> bool:
    if len(cards) < 5:
        return False

    record_rk_value = collections.defaultdict(int)
    for card in cards:
        rank = m.get_card_rank(card)
        rk_value = m.convert_rank_value(rank)
        record_rk_value[rk_value] += 1

    for rk_value, count in record_rk_value.items():
        if count >= 4:
            return True

    return False


def has_gold_full_house(cards: List[int], gold: int) -> bool:
    if len(cards) < 5:
        return False

    record_rk_value = collections.defaultdict(int)
    for card in cards:
        rank = m.get_card_rank(card)
        rk_value = m.convert_rank_value(rank)
        record_rk_value[rk_value] += 1

    check_gold = False
    check_pair = False
    for rk_value, count in record_rk_value.items():
        if rk_value == gold:
            if count == 3:
                check_gold = True
        else:
            if count >= 2:
                check_pair = True

    return check_gold and check_pair


def has_full_house(cards: List[int]) -> bool:
    if len(cards) < 5:
        return False

    record_rk_value = collections.defaultdict(int)
    for card in cards:
        rank = m.get_card_rank(card)
        rk_value = m.convert_rank_value(rank)
        record_rk_value[rk_value] += 1

    count_kind = 0
    count_pair = 0
    for rk_value, count in record_rk_value.items():
        if count >= 3:
            count_kind += 1
        else:
            if count == 2:
                count_pair += 1

    return count_kind >= 1 and (count_kind + count_pair >= 2)


def has_pair_with_flush(cards: List[int]) -> bool:
    if len(cards) < 5:
        return False

    def check_flush(poker: List[int]) -> bool:
        record_suit = collections.defaultdict(int)
        for p in poker:
            suit = m.get_card_suit(p)
            record_suit[suit] += 1

        for suit, count in record_suit.items():
            if count >= 3:
                return True
        return False

    record_pairs = collections.defaultdict(list)
    for card in cards:
        rank = m.get_card_rank(card)
        rk_value = m.convert_rank_value(rank)
        record_pairs[rk_value].append(card)

    for rk_value, pair in record_pairs.items():
        if len(pair) < 2:
            continue

        combinations = list(itertools.combinations(pair, 2))
        for c in combinations:
            values = []
            for card in cards:
                if card not in c:
                    values.append(card)

            if check_flush(values):
                return True

    return False


def has_flush(cards: List[int]) -> bool:
    if len(cards) < 5:
        return False

    record_suit = collections.defaultdict(int)
    for card in cards:
        suit = m.get_card_suit(card)
        record_suit[suit] += 1

    for suit, count in record_suit.items():
        if count >= 5:
            return True

    return False


def has_straight(cards: List[int]) -> bool:
    if len(cards) < 5:
        return False

    s = set()
    for card in cards:
        rank = m.get_card_rank(card)
        rk_value = m.convert_rank_value(rank)
        s.add(rk_value)

    if len(s) < 5:
        return False

    values = list(s)
    values.sort()
    cnt = 1
    pre_val = -1
    for x in values:
        if pre_val == -1:
            pre_val = x
            continue

        if x != pre_val + 1:
            cnt = 1
            pre_val = x
        else:
            cnt += 1
            pre_val = x

        if cnt >= 5:
            return True

    special_straight = [2, 3, 4, 5, 14]
    for ss in special_straight:
        if ss not in s:
            return False

    return True


def has_gold_three_of_a_kind(cards: List[int], gold: int) -> bool:
    if len(cards) < 3:
        return False

    record_rk_value = collections.defaultdict(int)
    for card in cards:
        rank = m.get_card_rank(card)
        rk_value = m.convert_rank_value(rank)
        record_rk_value[rk_value] += 1

    if record_rk_value[gold] >= 3:
        return True

    return False


def has_three_of_a_kind(cards: List[int]) -> bool:
    if len(cards) < 3:
        return False

    record_rk_value = collections.defaultdict(int)
    for card in cards:
        rank = m.get_card_rank(card)
        rk_value = m.convert_rank_value(rank)
        record_rk_value[rk_value] += 1

    for rk_value, count in record_rk_value.items():
        if count >= 3:
            return True

    return False


def has_two_pairs(cards: List[int]) -> bool:
    if len(cards) < 5:
        return False

    record_rk_value = collections.defaultdict(int)
    for card in cards:
        rank = m.get_card_rank(card)
        rk_value = m.convert_rank_value(rank)
        record_rk_value[rk_value] += 1

    pair_count = 0
    for rk_value, count in record_rk_value.items():
        if count >= 2:
            pair_count += 1

    if pair_count >= 2:
        return True
    return False


def has_one_pair(cards: List[int]) -> bool:
    if len(cards) < 2:
        return False

    record_rk_value = collections.defaultdict(int)
    for card in cards:
        rank = m.get_card_rank(card)
        rk_value = m.convert_rank_value(rank)
        record_rk_value[rk_value] += 1

    pair_count = 0
    for rk_value, count in record_rk_value.items():
        if count >= 2:
            pair_count += 1

    if pair_count >= 1:
        return True
    return False


def has_high_card(cards: List[int]) -> bool:
    if len(cards) < 1:
        return False
    return True
