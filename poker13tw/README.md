# poker13tw README

## 加一色13支 基本玩法說明
- 總共使用 65 張撲克牌，原 52 張多加一組黑桃。  
每局會切掉一張牌當作彩金牌(例如切掉方塊5，另外4張5都會變成彩金牌)  
- 每人會拿到 16 張牌，需組出 頭墩(3)/中墩(5)/尾墩(5) ，且牌型  
大小須滿足 頭墩 < 中墩 < 尾墩，而剩餘的 3 張牌不使用。  

- 先比牌型大小，牌型相同時只比數字不比花色。  

- 牌型大小: 見 const.go  

- 額外加分: 單一指定墩贏某家且滿足牌型，輸的玩家需支付額外分數  
詳見 getRecommendedCard.go  

## **pokerAPI.go**

### type PokerAPI interface   
``` go
type PokerAPI interface {

	// 給定一墩(只能是5張或3張)和彩金牌，返回該墩的分數
	GetPokerHandScore(cards []int, goldCard int) int

	// 給定整組手牌(16張)和彩金牌，返回推薦組法，格式為 頭墩/中墩/尾墩
	// 目前設定是先返回一組，可修改成返回多組
	GetRecommendPokerHand(cards []int, goldCard int) [][][]int

	// 返回一副新發的牌，格式為 彩金牌/玩家A手牌/玩家B手牌/玩家C手牌/玩家D手牌
	// 需要先設定亂數種子
	GetNewSplitDeck(rm *RandomManager) [][]int
}
```  

### uasge  
在 game server 啟動時先設置亂數種子  
rm := &poker13tw.RandomManager{}  
rm.SetRandSeed(time.Now().UnixNano())  

在每局中創建新的 PokerService  
pokerAPI := poker13tw.NewPokerService()

-----------------------------------------------------------
## **pokerService.go** 
用於定義 PokerService 類型，並提供了對應的 PokerAPI 方法的實現。


## **const.go**
用於定義遊戲中的不同牌型，並提供了一個基礎值 CombatBase 用於計算每墩的分數。通過將每墩分數除以 CombatBase，我們可以獲得對應的牌型。  
為遊戲的計分和分析提供了基礎，將不同的牌型轉換成數值，<span style="color:red;">還可用來幫助計算額外加分。</span>



## **convertTools.go**  
定義了撲克牌的相關轉換函式，使我們能夠在牌值、花色、順序和點數之間進行方便的互轉。根據牌值的不同範圍，我們將其映射到對應的花色、順序和點數，並提供了內部的轉換函式。

具體的定義如下：  
牌值（value）：範圍從 1 到 52，表示整副撲克牌中的每一張牌。  
花色（suit）：包括黑桃、愛心、方塊和梅花。  
順序（rank）：範圍從 2 到 A（代表 Ace）。  
點數（point）：範圍從 2 到 14，A 的點數為 14。  

### Note:   
01 ~ 13 為 黑桃 2,3,4...9 ,T, J, Q, K, A  
14 ~ 26 為 愛心 2,3,4...9 ,T, J, Q, K, A  
27 ~ 39 為 方塊 2,3,4...9 ,T, J, Q, K, A  
40 ~ 52 為 梅花 2,3,4...9 ,T, J, Q, K, A  


## **checkHandType.go**
包含用於評估不同牌型的函式，將一墩牌(只能是5張或3張)傳入並根據評估結果返回 true 或 false。  


## **calculateCombat.go**  
會將該墩計算得出對應的分數。  
分數格式如下: PokerHandType x CombatBase + offset  
先用牌型 x 一個極大的數。而對於相同牌型，則使用 offset 進行進一步區分。  
offset 是根據該牌型在比牌時的分組順序所得到的數字，每兩位數字一組。例如：  
fullHouse 77755 -> 777 + 55, offset = 0705  
twoPairs   66443 -> 66 + 44 +3, offset = 060403   
straight  34567 -> 34567, offset = 07  
### Note:  <span style="color:red;">需要搭配牌型判斷，才能使用。</span>  

## **checkCardCombination.go**
提供了檢查玩家手牌是否包含特定牌型的函式。
如果包含，則返回 true，否則返回 false。

## **getCardCombination.go**  
包含了函式，用於獲取不同牌型的核心牌組合。
### Note: 
1. 這些函式僅返回核心牌組，未必構成一墩牌。 例如鐵支 5555, 只會返回4張
2. 這些函式會返回所有可能的組合。例如順子 45678，其中可能包含梅花 4 與方塊 4，會返回兩個順子組合。另一個例子是順子 456789，會返回 45678 與 56789 兩個順子組合。
3. 返回的組合有經過排序，會按照 card 數值由大到小排序。
ex: 順子456789 會返回 [[9,8,7,6,5], [8,7,6,5,4]]

## **getRecommendedCard.go**
會根據先組頭墩(組3張) 或 中墩/尾墩(組5張) 來進行組牌。  

實作方式 order553 表示 先組5張牌 -> 再組5張牌 -> 再組3張牌。組牌時:  
1. 依照單一指定墩贏某家且滿足牌型的特殊牌型作為組牌優先順序。
2. 由大到小抽取各牌型的核心牌，3墩核心牌都組完後，再檢查將每墩牌補到足夠張數。  



## **deck.go**
定義了用於生成撲克牌牌組的函式，並包括隨機數管理相關功能。
RandomManager 類型是用於管理隨機數的結構，其中包含了 randSeed 屬性，用於存儲隨機種子。  
SetRandSeed 方法用於設置隨機種子，接受一個 int64 類型的種子值作為參數。  
GetRandSeed 方法用於獲取當前隨機種子，返回一個 int64 類型的種子值。  
GetNewDeck 函式用於生成一副新的撲克牌牌組，接受一個 RandomManager 的指針作為參數。它通過隨機數生成牌組並進行洗牌操作，然後返回一副新的撲克牌牌組，做為該局遊戲使用。  

## **cardApplication.go**  

用 python 撰寫初版時的 API 函式，後轉換成使用Go語言的interface來實現外部接口。  

## **auxiliary.go**
用 python 撰寫初版時的內建函式，後轉換成使用 Go 語言時實作了這些 python 的內建函式。  



