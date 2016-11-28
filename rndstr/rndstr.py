# -*- coding: utf-8 -*-
import random
import datetime

def RandomCharacter(count):
    chars = []
    for i in range(0, count):
        if random.randint(0, 1)==0:
            chars.append(unichr(random.randint(33, 126)))
        else:
            chars.append(unichr(random.randint(0x4E00, 0x9FA5)))

    return ''.join(chars)


if __name__=="__main__":
    start = datetime.datetime.now()
    RandomCharacter(2000000)
    end = datetime.datetime.now()
    # 0:00:07.423000
    print end-start
