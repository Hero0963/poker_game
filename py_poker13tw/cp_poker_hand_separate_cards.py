import cp_poker_hand_combinations as cbn
import cp_const as c
import cp_check_poker_hand as m
from typing import List
import collections
import copy
import itertools

"""
組牌邏輯:玩家手中16張牌，從最大牌型開始判斷回傳所有組法
"""


def get_five_of_a_kind_from_poker_hand(cards: List[int]):
    if not cbn.has_five_of_a_kind(cards[:]):
        return []
    result = []
    record_rk_value = collections.defaultdict(int)
    record_split_cards = collections.defaultdict(list)
    for card in cards:
        rank = m.get_card_rank(card)
        rk_value = m.convert_rank_value(rank)
        record_rk_value[rk_value] += 1
        record_split_cards[rk_value].append(card)

    for rk_value, count in record_rk_value.items():
        if count >= 5:
            s = record_split_cards[rk_value]
            result.append(s)

    return result


def recover_original_point_by_rk_value_and_suit(rk_value: int, suit_recover: int) -> int:
    return rk_value - 1 + 13 * suit_recover


def get_straight_flush_from_poker_hand(cards: List[int]):
    if not cbn.has_straight_flush(cards):
        return []

    result = []

    def get_straight_flush(poker: List[int]):
        res = []
        if len(poker) < 5:
            return res

        suit_recover = poker[0] // 13
        points = []
        for p in poker:
            rank = m.get_card_rank(p)
            rk_value = m.convert_rank_value(rank)
            points.append(rk_value)

        points.sort()

        record = []
        pre_val = -1
        for p in points:
            if pre_val == -1:
                record = [p]
                pre_val = p
                continue

            if p != pre_val + 1:
                record = [p]
            else:
                record.append(p)
                pre_val = p

            if len(record) == 5:
                res.append(copy.deepcopy(record))
                record.pop(0)

        special_straight = [2, 3, 4, 5, 14]
        check_special_straight = True
        sp = set(points)
        for ss in special_straight:
            if ss not in sp:
                check_special_straight = False
                break

        if check_special_straight:
            res.append(special_straight)

        ans = []
        for rk_values in res:
            straight_flush = []
            for v in rk_values:
                r = recover_original_point_by_rk_value_and_suit(v, suit_recover)
                straight_flush.append(r)

            ans.append(copy.deepcopy(straight_flush))

        return ans

    split_by_suit = collections.defaultdict(list)
    for card in cards:
        suit = m.get_card_suit(card)
        split_by_suit[suit].append(card)

    for _, single_suit_cards in split_by_suit.items():
        res = get_straight_flush(single_suit_cards)
        if res:
            for r in res:
                result.append(r)

    return result


def get_five_of_gold_tiger_from_poker_hand(cards: List[int], gold: int):
    if not cbn.has_five_of_gold_tiger(cards, gold):
        return []

    result = []
    res = []

    for card in cards:
        rank = m.get_card_rank(card)
        rk_value = m.convert_rank_value(rank)
        if rk_value == gold:
            res.append(card)

    result.append(res)
    return result


def get_four_of_a_kind_from_poker_hand(cards: List[int]):
    if not cbn.has_four_of_a_kind(cards):
        return []

    result = []
    split_by_rk_value = collections.defaultdict(list)
    for card in cards:
        rank = m.get_card_rank(card)
        rk_value = m.convert_rank_value(rank)
        split_by_rk_value[rk_value].append(card)

    for _, srv in split_by_rk_value.items():
        if len(srv) >= 4:
            all_four_kind = list(itertools.combinations(srv, 4))
            for kind in all_four_kind:
                res = []
                for kd in kind:
                    res.append(kd)
                result.append(res)

    return result


def get_gold_full_house_from_poker_hand(cards: List[int], gold: int):
    if not cbn.has_gold_full_house(cards[:], gold):
        return []

    result = []
    golden_card = []
    record_rk_values = collections.defaultdict(list)
    for card in cards:
        rank = m.get_card_rank(card)
        rk_value = m.convert_rank_value(rank)
        if rk_value == gold:
            golden_card.append(card)
        else:
            record_rk_values[rk_value].append(card)

    golden_all_possible = list(itertools.combinations(golden_card, 3))
    for golden in golden_all_possible:
        for _, values in record_rk_values.items():
            if len(values) >= 2:
                pairs = list(itertools.combinations(values, 2))
                for p in pairs:
                    res = []
                    for g in golden:
                        res.append(g)
                    for pp in p:
                        res.append(pp)
                    result.append(res)

    return result


def get_full_house_from_poker_hand(cards: List[int]):
    if not cbn.has_full_house(cards[:]):
        return []
    result = []
    record_rk_values = collections.defaultdict(list)
    for card in cards:
        rank = m.get_card_rank(card)
        rk_value = m.convert_rank_value(rank)
        record_rk_values[rk_value].append(card)

    all_kinds = []
    all_pairs = []
    for _, values in record_rk_values.items():
        if len(values) >= 3:
            kinds = list(itertools.combinations(values, 3))
            for kind in kinds:
                all_kinds.append(kind)

        if len(values) >= 2:
            pairs = list(itertools.combinations(values, 2))
            for pair in pairs:
                all_pairs.append(pair)

    for kind in all_kinds:
        rk_value_kind = m.convert_rank_value(m.get_card_rank(kind[0]))
        for pair in all_pairs:
            rk_value_pair = m.convert_rank_value(m.get_card_rank(pair[0]))
            if rk_value_kind != rk_value_pair:
                res = []
                for kd in kind:
                    res.append(kd)
                for pr in pair:
                    res.append(pr)

                result.append(res)

    return result


def get_pair_with_flush_from_poker_hand(cards: List[int]):
    if not cbn.has_pair_with_flush(cards[:]):
        return []

    result = []

    record_suit = collections.defaultdict(list)
    for card in cards:
        suit = m.get_card_suit(card)
        record_suit[suit].append(card)

    for suit, values in record_suit.items():
        flushes = []
        pair_card = -1
        for card in values:
            if card not in flushes:
                flushes.append(card)
            else:
                pair_card = card
                flushes.remove(card)

        if pair_card != -1:
            flush = list(itertools.combinations(flushes, 3))
            for fh in flush:
                res = [pair_card, pair_card]
                for f in fh:
                    res.append(f)
                result.append(res)

    return result


def get_flush_from_poker_hand(cards: List[int]):
    if not cbn.has_flush(cards[:]):
        return []

    result = []
    record_suit = collections.defaultdict(list)
    for card in cards:
        suit = m.get_card_suit(card)
        record_suit[suit].append(card)

    for suit, values in record_suit.items():
        if len(values) >= 5:
            suit_all_possibles = list(itertools.combinations(values, 5))
            for suit_ap in suit_all_possibles:
                res = []
                for sap in suit_ap:
                    res.append(sap)
                result.append(res)
    return result


def get_straight_from_poker_hand(cards: List[int]):
    if not cbn.has_straight(cards[:]):
        return []
    result = []

    points = [set() for _ in range(14 + 1)]
    for card in cards:
        rank = m.get_card_rank(card)
        rk_value = m.convert_rank_value(rank)
        points[rk_value].add(card)

    cnt = 0
    prev = -1
    straights = []
    for i in range(1, len(points)):
        if not points[i]:
            cnt = 0
            prev = -1

        else:
            if prev == -1:
                cnt = 1
                prev = i
            else:
                if prev != i - 1:
                    cnt = 1
                    prev = i
                else:
                    cnt += 1
                    prev = i

        if cnt >= 5:
            straights.append([i - 4, i - 3, i - 2, i - 1, i])

    check_special_straight = True
    special_straight = [2, 3, 4, 5, 14]
    for ss in special_straight:
        if not points[ss]:
            check_special_straight = False
            break

    if check_special_straight:
        straights.append(special_straight)

    for ss in straights:
        cs = []
        for s in ss:
            cs.append(points[s])
        all_possibles = list(itertools.product(*cs))
        for ap in all_possibles:
            res = []
            for p in ap:
                res.append(p)
            result.append(res)

    return result


def get_gold_three_of_a_kind_from_poker_hand(cards: List[int], gold: int):
    if not cbn.has_gold_three_of_a_kind(cards[:], gold):
        return []
    result = []
    golden = []
    for card in cards:
        rank = m.get_card_rank(card)
        rk_value = m.convert_rank_value(rank)
        if rk_value == gold:
            golden.append(card)

    all_possibles = list(itertools.combinations(golden, 3))
    for ap in all_possibles:
        result.append(list(ap))

    return result


def get_three_of_a_kind_from_poker_hand(cards: List[int]):
    if not cbn.has_three_of_a_kind(cards[:]):
        return []

    result = []
    record_rk_values = collections.defaultdict(list)

    for card in cards:
        rank = m.get_card_rank(card)
        rk_value = m.convert_rank_value(rank)
        record_rk_values[rk_value].append(card)

    for _, poker in record_rk_values.items():
        if len(poker) >= 3:
            all_possibles = list(itertools.combinations(poker, 3))
            for ap in all_possibles:
                result.append(list(ap))

    return result


def get_two_pairs_from_poker_hand(cards: List[int]):
    if not cbn.has_two_pairs(cards[:]):
        return []

    result = []
    record_rk_values = collections.defaultdict(list)

    for card in cards:
        rank = m.get_card_rank(card)
        rk_value = m.convert_rank_value(rank)
        record_rk_values[rk_value].append(card)

    pairs = []
    for rk_value, poker in record_rk_values.items():
        if len(poker) >= 2:
            pairs.append(poker)

    two_pairs = list(itertools.combinations(pairs, 2))
    for two_pair in two_pairs:
        tp_0 = list(itertools.combinations(two_pair[0], 2))
        tp_1 = list(itertools.combinations(two_pair[1], 2))
        for tp0 in tp_0:
            for tp1 in tp_1:
                res = list(set(tp0) | set(tp1))
                result.append(res)

    return result


def get_one_pair_from_poker_hand(cards: List[int]):
    if not cbn.has_one_pair(cards[:]):
        return []

    result = []
    record_rk_values = collections.defaultdict(list)

    for card in cards:
        rank = m.get_card_rank(card)
        rk_value = m.convert_rank_value(rank)
        record_rk_values[rk_value].append(card)

    pairs = []
    for rk_value, poker in record_rk_values.items():
        if len(poker) >= 2:
            pairs.append(poker)

    for pair in pairs:
        all_possible = list(itertools.combinations(pair, 2))
        for ap in all_possible:
            result.append(list(ap))

    return result


def get_high_card_from_poker_hand(cards: List[int]):
    if not cbn.has_high_card(cards[:]):
        return []

    result = []
    for card in cards:
        result.append([card])
    return result
