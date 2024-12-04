# ChatBot阶段汇报22.04

## 技术栈：python+tensorflow2+Keras

原技术栈： python+pytorch+tensor_gpu1.x   

更换原因：因为hpc的显卡驱动版本太新了，pytorch不兼容（一些模型超过一年没人更新了）；另一方面pytorch和前后端的兼容性不太好，不如tf。

目前版本：cuda11.6+python3.7+tensorflow2.7+Keras

> 不同版本的tf已经创建了conda环境，如果想用其他版本新建一个conda env，不要更新已有conda的包版本！
> 
> 已有环境：
> 
> > /home/hpcuser/.conda/envs/tf1.15                 %tf-idf 环境
> > /home/hpcuser/.conda/envs/sence-graph
> > /home/hpcuser/.conda/envs/sgg
> > /home/hpcuser/.conda/envs/sgg_bm
> > /home/hpcuser/.conda/envs/torch04
> > /home/hpcuser/.conda/envs/py3_tf1.15           %tf-idf 环境
> > /opt/anaconda
> > /opt/anaconda/envs/unsup_seg
> > /home/hpcuser/.conda/envs/myconda
> > /opt/anaconda/envs/chatbox
> > /opt/anaconda/envs/deepmd-kit
> > /opt/anaconda/envs/tmp
> > /opt/anaconda/envs/zyx
> > /opt/anaconda/envs/py3_tf2.6                                  %废了，没删
> > /home/hpcuser/anaconda3
> > /home/hpcuser/anaconda3/envs/chat                    %废了，没删
> > /home/hpcuser/anaconda3/envs/Bert                    %gpu无法识别
> > /home/hpcuser/anaconda3/envs/Bert_tf2              %tf2.0 （gpu无法识别)
> > /home/hpcuser/anaconda3/envs/tf2.7                    %现使用环境

> source activate tf2.7

使用环境前检查gpu是否可用，部分环境gpu不可用

> print(tf.test.is_gpu_available())

## 数据集的构建

上个学期，我们由爬虫得到的一整年度的qq聊天记录，洗出了200+条相关问答文档，相关数据库内容已整理至服务器/home/hpcuser/Chatbot/data/

qs.txt文件是询问集，开头以[qry]标记。

ps.txt文件是回答集，开头以[ans]标记。

qq.txt是两者结合，每行包括对应的询问和回答。

具体的建立数据过程中的语法统一和标点符号去歧义规则，见Chatbot的git页面。

**TODO：**

1. 将问答写的再官方一点，现有数据库有很多“是的”“不行”之类的回答，问句也有些过于简短，对模型不好识别。

   2. 根据测试集结果，补充常见问题的答复。

3. 思路上的问题：对于专业的名词和报错、网址、代码，没有找到合适的tokenization和encoding方法。

## 前序工作：tf-idf版本的实现

此前，我因为没有找到合适的现成模型框架，根据if-idf算法在tf1.15版本下手敲了一个粗糙的代码实现此功能，也顺便探路了一下整个工作的流程。手敲代码可以跑起来且很快，但准确率堪忧。大家如果想要尝试，可以使用/home/hpcuser/Chatbot/tf-idf.py。

测试集使用的是/home/hpcuser/Chatbot/data/test_q.txt，可自行修改其中的内容。

在实践过程中，我首先确定了tokenization可以使用jieba，他的准确性非常有保障（人肉核验），此外，该库支持自定义词典，且不允许不添加词频，对于专有词汇例如：

> 创新港
> 兴庆校区
> 曲江校区
> 雁塔校区
> 移动交通大学

我将其放在userdict.txt中.

此外,停词表我也已经找到了现成的/home/hpcuser/Chatbot/stop_words.txt,该文档是川大整理的,感觉够用了.

## 实现思路：Squad的实现

[The Stanford Question Answering Dataset](https://rajpurkar.github.io/SQuAD-explorer/#:~:text=Stanford%20Question%20Answering%20Dataset%20%28SQuAD%29%20is%20a%20reading,reading%20passage%2C%20or%20the%20question%20might%20be%20unanswerable.)[The Stanford Question Answering Dataset](https://rajpurkar.github.io/SQuAD-explorer/#:~:text=Stanford%20Question%20Answering%20Dataset%20%28SQuAD%29%20is%20a%20reading,reading%20passage%2C%20or%20the%20question%20might%20be%20unanswerable.)

Squad是一个问答型的数据库，数据集内容如下：

![](C:\Users\ZachW\AppData\Roaming\marktext\images\2022-04-09-14-40-51-image.png)

目前我们使用的是squad2.0来对我们的模型进行检验。

数据格式大致如下：

> dta[0]
> 
> # 
> 
> {'title': 'Beyoncé',
>  'paragraphs': [
>                  {'qas': [{'question': 'When did Beyonce start becoming popular?',
>                                  'id': '56be85543aeaaa14008c9063',
>                          'answers': [{'text': 'in the late 1990s', 'answer_start': 269}],
>                       'is_impossible': False}]],
>                              'context':'Beyoncé Giselle Knowles-Carter (/biːˈjɒnseɪ/ bee-YON-say) (born September 4, 1981) is an American singer, songwriter, record producer and actress. Born and raised in Houston, Texas, she performed in various singing and dancing competitions as a child, and rose to fame in the late 1990s as lead singer of R&B girl-group Destiny\'s Child. Managed by her father, Mathew Knowles, the group became one of the world\'s best-selling girl groups of all time. Their hiatus saw the release of Beyoncé\'s debut album, Dangerously in Love (2003), which established her as a solo artist worldwide, earned five Grammy Awards and featured the Billboard Hot 100 number-one singles "Crazy in Love" and "Baby Boy".'}
>      }

我们主要使用的是qas部分，question为询问，answers为答复。

Squad数据集已整理在/home/hpcuser/Chatbot/BERT/SQuAD/内，里面包括v1.0和v2.0，两者可以通用，但是v2.0更新些，数据中也有参数标记每句话的版本。

具体的接口形式可以参考，/home/hpcuser/Chatbot/BERT/bert-master/run_squad.py。

在运行后，我们把我们的回答集运行squad的evalute.py，可以得到正确率等结果参数，作为模型调试的结论。

## BERT模型

这个展开来可以说很多,现有的squad榜单大部分都是基于bert做的,新一点的有albert等等.大家需要搞懂基础的bert原理,首先需要看paper和code

https://github.com/google-research/bert

相关的中英文预训练模型也已放在了BERT/文件夹中,英文encode推荐使用uncased模型,中文还未测试过.

看代码的时候请注意,其中内容为tf1.x版本的,许多语法已经不再使用,我们需要理解的仅是思路的结构即可.

具体实现由于我们的显卡驱动太新了,我们只能使用最新版本的tf2.7与Keras库,具体学习可参考官方doc,其他二手整理资料现在还比较少.

[Module: tf &nbsp;|&nbsp; TensorFlow Core v2.8.0](https://tensorflow.google.cn/api_docs/python/tf)

[Keras API reference](https://keras.io/api)

其他资料我可以后续补充.

**#TODO:**

**将tensorflow1.x的bert改成支持Keras.**
