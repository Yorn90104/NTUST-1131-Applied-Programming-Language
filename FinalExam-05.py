class SortedMap():
    def __init__(self):
        self.data = dict()
    
    def Exists(self, key) -> bool:
        return key in self.data
    
    def Add(self, key: any, value: any):
        if self.Exists(key):
            raise KeyError(f"此鍵已存在: {key}")
        else:
            self.data[key] = value

    def Remove(self, key):
        if self.Exists(key):
            del self.data[key]
        else:
            raise KeyError(f"此鍵不存在: {key}")
            
    def Keys(self) -> list:
        return sorted(self.data.keys())
    
    def Find(self, key) -> any:
        if self.Exists(key):
            return self.data[key]
        else:
            raise KeyError(f"此鍵不存在: {key}")
        
if __name__ == "__main__":
    # 創建 SortedMap 實例
    sm = SortedMap()

    # 添加鍵值對
    sm.Add("b", 2)
    sm.Add("a", 1)
    sm.Add("c", 3)

    # 嘗試添加已存在的鍵
    try:
        sm.Add("a", 99)
    except KeyError as e:
        print(e)

    # 檢查鍵是否存在
    print(sm.Exists("b"))  # 輸出：True
    print(sm.Exists("z"))  # 輸出：False

    # 查找鍵的值
    print(sm.Find("a"))    # 輸出：1

    # 嘗試查找不存在的鍵
    try:
        print(sm.Find("z"))
    except KeyError as e:
        print(e)

    # 移除鍵
    sm.Remove("b")
    print(sm.Exists("b"))  # 輸出：False

    # 嘗試移除不存在的鍵
    try:
        sm.Remove("z")
    except KeyError as e:
        print(e)

    # 獲取排序後的鍵、值和鍵值對
    print(sm.Keys())       # 輸出：['a', 'c']
