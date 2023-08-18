from dataclasses import dataclass

"""
基本牌型大小:
五枚 > 同花順 > 彩金五虎將 > 鐵支 > 彩金葫蘆 >
對子同花 > 同花 > 順子 > 彩金三條 > 三條 >兩對　> 一對 > 單張

五枚 (Five of a Kind): 五張點數相同的牌
同花順 (Straight Flush): 相同花色的順子
彩金五虎將 (Five of Gold Tiger): 四張彩金牌 + 一張雜牌
鐵支 (Four of a Kind): 四張點數相同的牌外加任一單張
彩金葫蘆 (Gold Full House): 三張彩金牌配上一對
葫蘆 (Full House): 一組對子加上一組三條
對子同花 (Pair with Flush): 兩張相同點數與花色的牌，外加三張相同花色的牌
同花 (Flush): 五張相同花色的牌
順子 (Straight): 五張點數各差一點的連續牌
彩金三條 (Gold Three of a Kind): 三張彩金牌
三條 (Three of a Kind): 三張同點數的牌
兩對 (Two Pairs): 兩組對子，外加任一單張
一對 (One Pair): 兩張同點數的牌
單張(烏龍) (High Card): 單張不成牌，又稱烏龍
"""


@dataclass
class PokerHand:
    FIVE_HIGH: int = 14
    STRAIGHT_FLUSH: int = 13
    FIVE_OF_GOLD_TIGER: int = 12
    FOUR_OF_A_KIND: int = 11
    GOLD_FULL_HOUSE: int = 10
    FULL_HOUSE: int = 9
    PAIR_WITH_FLUSH: int = 8
    FLUSH: int = 7
    STRAIGHT: int = 6
    GOLD_THREE_OF_A_KIND: int = 5
    THREE_OF_A_KIND: int = 4
    TWO_PAIRS: int = 3
    ONE_PAIR: int = 2
    HIGH_CARD: int = 1


COMBAT_BASE = 1_00_00_00_00_00_00_00_00
