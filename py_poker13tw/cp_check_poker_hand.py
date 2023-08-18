import cp_const as c
import random
from typing import List
import collections

PH = c.PokerHand()

"""
生出一副牌
"""


class PokerDeck:
    def __init__(self):
        self.deck = [i for i in range(1, 52 + 1)] + [j for j in range(1, 13 + 1)]

    def shuffle_deck(self):
        random.shuffle(self.deck)


"""
1~52 與 點數花色 互轉
"""


def get_card_suit(card):
    suits = ['黑桃', '愛心', '方塊', '梅花']
    return suits[(card - 1) // 13]


def get_card_rank(card):
    ranks = ['2', '3', '4', '5', '6', '7', '8', '9', '10', 'J', 'Q', 'K', 'A']
    return ranks[(card - 1) % 13]


def convert_to_value(suit, rank):
    rank -= 1
    if rank == 0:
        rank = 13
    if suit == "黑桃":
        return rank
    elif suit == "愛心":
        return rank + 13
    elif suit == "方塊":
        return rank + 26
    elif suit == "梅花":
        return rank + 39
    else:
        raise ValueError("無效的花色")


def convert_rank_value(rank):
    rk = ['', '', '2', '3', '4', '5', '6', '7', '8', '9', '10', 'J', 'Q', 'K', 'A']
    return rk.index(rank)


"""
判斷合法牌
"""


# 待補
def is_legal_cards(cards: List[int]) -> bool:
    # 待補
    return True


"""
每種牌型的判斷
"""


def is_five_of_a_kind(cards: List[int]) -> bool:
    if len(cards) != 5:
        return False

    record_rk_value = collections.defaultdict(int)
    for card in cards:
        suit, rank = get_card_suit(card), get_card_rank(card)
        rk_value = convert_rank_value(rank)
        record_rk_value[rk_value] += 1

    for rk_value, count in record_rk_value.items():
        if count == 5:
            return True
    return False


def is_straight_flush(cards: List[int]) -> bool:
    if len(cards) != 5:
        return False

    record_suit = collections.defaultdict(int)
    record_rk_value = []
    for card in cards:
        suit, rank = get_card_suit(card), get_card_rank(card)
        record_suit[suit] += 1
        rk_value = convert_rank_value(rank)
        record_rk_value.append(rk_value)

    if len(record_suit) > 1:
        return False

    record_rk_value.sort()
    cnt = 1
    pre_val = -1
    for x in record_rk_value:
        if pre_val == -1:
            pre_val = x
            continue

        if x != pre_val + 1:
            break
        else:
            cnt += 1
            pre_val = x
    if cnt == 5:
        return True

    if cnt == 4 and record_rk_value[0] == 2 and record_rk_value[-1] == 14:
        return True

    return False


def is_five_of_gold_tiger(cards: List[int], gold: int) -> bool:
    if len(cards) != 5:
        return False

    record_rk_value = collections.defaultdict(int)
    for card in cards:
        suit, rank = get_card_suit(card), get_card_rank(card)
        rk_value = convert_rank_value(rank)
        record_rk_value[rk_value] += 1

    if record_rk_value[gold] >= 4:
        return True

    return False


def is_four_of_a_kind(cards: List[int]) -> bool:
    if len(cards) != 5:
        return False

    record_rk_value = collections.defaultdict(int)
    for card in cards:
        suit, rank = get_card_suit(card), get_card_rank(card)
        rk_value = convert_rank_value(rank)
        record_rk_value[rk_value] += 1

    for rk_value, count in record_rk_value.items():
        if count >= 4:
            return True

    return False


def is_gold_full_house(cards: List[int], gold: int) -> bool:
    if len(cards) != 5:
        return False

    record_rk_value = collections.defaultdict(int)
    for card in cards:
        suit, rank = get_card_suit(card), get_card_rank(card)
        rk_value = convert_rank_value(rank)
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


def is_full_house(cards: List[int]) -> bool:
    if len(cards) != 5:
        return False

    record_rk_value = collections.defaultdict(int)
    for card in cards:
        suit, rank = get_card_suit(card), get_card_rank(card)
        rk_value = convert_rank_value(rank)
        record_rk_value[rk_value] += 1

    check_three = False
    check_pair = False
    for rk_value, count in record_rk_value.items():
        if count == 3:
            check_three = True
        if count == 2:
            check_pair = True

    return check_three and check_pair


def is_pair_with_flush(cards: List[int]) -> bool:
    if len(cards) != 5:
        return False

    record_suit = collections.defaultdict(int)
    for card in cards:
        suit = get_card_suit(card)
        record_suit[suit] += 1

    if len(record_suit) != 1:
        return False

    if len(set(cards)) == 5:
        return False

    return True


def is_flush(cards: List[int]) -> bool:
    if len(cards) != 5:
        return False

    record_suit = collections.defaultdict(int)
    for card in cards:
        suit, rank = get_card_suit(card), get_card_rank(card)
        record_suit[suit] += 1

    if len(record_suit) == 1:
        return True
    return False


def is_straight(cards: List[int]) -> bool:
    if len(cards) != 5:
        return False

    record_rk_value = []
    for card in cards:
        suit, rank = get_card_suit(card), get_card_rank(card)
        rk_value = convert_rank_value(rank)
        record_rk_value.append(rk_value)

    record_rk_value.sort()
    cnt = 1
    pre_val = -1
    for x in record_rk_value:
        if pre_val == -1:
            pre_val = x
            continue

        if x != pre_val + 1:
            break
        else:
            cnt += 1
            pre_val = x
    if cnt == 5:
        return True

    if cnt == 4 and record_rk_value[0] == 2 and record_rk_value[-1] == 14:
        return True
    return False


def is_gold_three_of_a_kind(cards: List[int], gold: int) -> bool:
    record_rk_value = collections.defaultdict(int)
    for card in cards:
        rank = get_card_rank(card)
        rk_value = convert_rank_value(rank)
        record_rk_value[rk_value] += 1

    if record_rk_value[gold] >= 3:
        return True

    return False


def is_three_of_a_kind(cards: List[int]) -> bool:
    record_rk_value = collections.defaultdict(int)
    for card in cards:
        rank = get_card_rank(card)
        rk_value = convert_rank_value(rank)
        record_rk_value[rk_value] += 1

    for rk_value, count in record_rk_value.items():
        if count >= 3:
            return True

    return False


def is_two_pairs(cards: List[int]) -> bool:
    if len(cards) != 5:
        return False

    record_rk_value = collections.defaultdict(int)
    for card in cards:
        rank = get_card_rank(card)
        rk_value = convert_rank_value(rank)
        record_rk_value[rk_value] += 1

    pair_count = 0
    for rk_value, count in record_rk_value.items():
        if count >= 2:
            pair_count += 1

    if pair_count >= 2:
        return True
    return False


def is_one_pair(cards: List[int]) -> bool:
    record_rk_value = collections.defaultdict(int)
    for card in cards:
        rank = get_card_rank(card)
        rk_value = convert_rank_value(rank)
        record_rk_value[rk_value] += 1

    pair_count = 0
    for rk_value, count in record_rk_value.items():
        if count >= 2:
            pair_count += 1

    if pair_count >= 1:
        return True
    return False


def is_high_card(cards: List[int]) -> bool:
    if len(cards) < 1:
        return False
    return True
