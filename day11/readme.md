有点意思的题目，记录一下思路:

在 part2 中，题目中有说:

> You're worried you might not ever get your items back. So worried, in fact, that your relief that a monkey's inspection didn't damage an item no longer causes your worry level to be divided by three.

> Unfortunately, that relief was all that was keeping your worry levels from reaching ridiculous levels. You'll need to find another way to keep your worry levels manageable.

> At this rate ...

这时候题目就给了暗示说，Part 1 中的 `worryLevel / 3` 不再适用了，(其实后来想想确实不适用了，因为 `rounds` 变多的情况下，指数级别的增长就有很大可能爆掉 `uint64` 了)。

所以 Part 2 的大意就可以翻译为: 你自己自定义一个 rate 来保证题目结果(或者说迭代?)不变的情况下，控制 `worryLevel` 的增长，使其在 round 10000 时 `worryLevel` 不会溢出。

(看Reddit上的讨论，大部分人跟看题干一个小时之内的我一样并没有领会到这层含义，一直在 debug 迭代过程觉得自己哪里写错了)

过了这坎之后就要开始思考怎么去计算这个 rate 了，再次梳理一下这个`Monkey in the middle` 游戏的过程:

0. 该死的猴子抢走了你的东西，你怕猴子把你宝贵的~~老婆们~~给整坏了,所以就有了 `worryLevel` 的定义。
1. 第一只猴子看的时候，你对你~~老婆们~~的担心程度是 n (假设有两个)，第一只猴子看完(inspect过程)之后，你的担忧程度会增加(* or +)，这时候猴子看你的担心程度来扔给其它兄弟(比如说 `n % 19 ==0 ` 时扔给第二只猴子), 以此往复,计算出每只猴子经手你~~老婆们~~的次数。
2. 所以，控制这个迭代过程不变的关键就变成了: **只要保证每次猴子经手顺序不乱，最后的结果就是一定的**，并不关抽象出来的`worryLevel`什么事。

这样问题就被简化了，我们只要找到猴子怎么决定经手顺序就可以了，没错，就是猴子根据你的担心程度调戏你的过程(也就是题目中的`Test`过程)，他们心里也有一把秤来衡量这东西应该交给哪个兄弟，这个**秤的刻度**就是每次`Test`过程 mod 的值。

所以，如果我们能重新定义猴子心里的秤的刻度就可以保证他们的经手顺序不乱了，再简化下模型:

```
有4只猴子，每只猴子的秤刻度分别是 3, 5, 7, 11 (记为 a1, a2, a3, a4)，
计算出一个正整数 x 使得 (worryLevel op x) % a == worryLevel % a (a = a1, a2, a3, a4), 另外要使得 worryLevel op x << worryLevel
```


这里的 op x 就是最后我们想要的 `rate` 的解。

根据 [mod 运算符的计算规则](https://libraryguides.centennialcollege.ca/c.php?g=717548&p=5121841):

我们可以得到可以使用的 `rate` 结果为:

* 如果 op 是 - , 我们可以得到 x % a == 0，所以 x == 0 或者 x 同时是 (a1, a2, a3, a4) 的倍数，也就是 最小公倍数(Least Common Multiple)。

* 如果 op 是 / , x 只能是 1， 并且并不满足 < worryLevel 的条件。

* 但如果 op 本身就是 mod 呢 ？ mod 操作本身就可以有效控制`worryLevel`的无限制增长，只要 x 还是满足同时是 (a1, a2, a3, a4) 的最小公倍数就可以了。

综上， op 是 mod 的解优于 op 是 - 的解，所以最后的 `rate` 为: `worryLevel mod (multiply (a...))`。

妙啊，太妙了！
