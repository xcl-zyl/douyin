### Go学习
<br>     

#### 学习资料
1. [Go 语言上手 - 基础语言](https://juejin.cn/post/7093721879462019102)
2. [Go语言圣经](https://books.studygolang.com/gopl-zh/)
3. [青训营完整手册](https://bytedance.feishu.cn/docs/doccnFRB1TXYJPK6yprPETHLXgd#q8ZYps)
4. [抖音项目-青训营](https://bytedance.feishu.cn/docx/doxcnbgkMy2J0Y3E6ihqrvtHXPg)
5. [Go语言上手PPT](https://bytedance.feishu.cn/file/boxcnQnHXuDOdzd8CqVid7nQLmg)
6. [课后项目实践作业](https://juejin.cn/post/7094452391101071367/)
7. [Go语言工程实践PPT](https://bytedance.feishu.cn/file/boxcnRmlw9MjbtAMBnOW44y8dZd?hash=7cfc75acc80372c08463b622df90a4b5)
8. [高质量与性能调优PPT](https://bytedance.feishu.cn/file/boxcnqqWtT0xgWAIMGWVs7wM6fd?hash=ab6bfba21a54c52073c7341ecb3ab470)

<br>

#### 实践
1. 视频Feed流、视频投稿、个人信息：    

&emsp;&emsp;支持所有用户刷抖音，按投稿时间倒叙推出，登录用户可自己拍视频投稿，查看自己的基本信息和投稿列表，注册用户流程简化。
   - 支持所有用户刷抖音，按投稿时间倒叙推出，单次最多30个视频
   - 注册登录功能  
   - 实现登录用户投稿功能
   - 查看自己基本信息和投稿列表

根据分析，第一步应该设计数据库的表结构：
   - 首先实现基础功能：
   - 应包含用户表、视频表

**第一部分已全部完成**   
*存在问题： 1.只有重启程序才能刷新首页，没有刷新功能 2.封面图没有上传*
   
2. 点赞列表，用户评论
- 对视频点赞
- 对视频评论
- 个人主页有视频点赞列表

实现：
- 点赞功能：
  - 点赞时，将视频id加入用户喜爱表中，如果喜爱表不存在则建立喜爱表，显示个人喜欢页面时，读取喜爱表，返回喜爱列表视频，同时修改视频点赞数量，需要写入数据库
  - 取消赞时，将视频id从喜爱表中去掉，同时减少视频点赞数量，需要写入数据库
  - 用户登录之后，每个显示视频初始化时，判断这个视频在不在该用户喜爱表中，从而决定是否显示爱心
- 评论功能：
  - 评论功能相对简单一些，评论时将评论放入到评论列表显示到界面上，也应该根据视频id建立对应评论表，将新的评论写入数据库
完成点赞和评论功能后，应修改视频初始化，需要读取是否点赞，评论数，以及点赞数
  - 在完成这部分功能时，需要在数据库video表中增加内容，同时点赞和评论需要动态增加用户的喜爱视频表和对应视频的评论内容表，    
  - 表名建议为favorite_用户id，comment_视频id

3. 关注列表，粉丝列表
- 关注功能
  - 和点赞功能类似， 点击+关注时，将视频的发布用户加入登录用户的关注表中，没有关注表则动态建立关注表，同时修改用户关注数，和对应用户粉丝数，将登录用户加入到发布视频用户的粉丝表中，没有粉丝表则动态生成粉丝表，修改发布用户粉丝数
- 关注数和粉丝数加入用户信息中
  - 数据库的user表需要添加关注数和粉丝数
- 能够显示关注列表和粉丝列表
  - 显示关注列表和粉丝列表只需要读数据库对应表就行
  需要注意的是初始化视频显示时，需判断“+”即关注符号需不需要显示，和点赞功能类似，读取登录用户的关注列表判断是否应该显示“+”
  - 表名规则同上

<br>

  &emsp;&emsp;初始数据库文件在sql文件夹中，运行程序前需先配置好数据库环境，将数据导入数据库中 [数据库](sql/douyin.sql)

  **提示：连接数据库时记得修改代码中的数据库账户密码，在tool文件中**

## （andy）测试中--
## 源码

### sql文档

1.更改utf8mb4_0900_ai_ci为utf8mb4_general_ci以解决数据库版本问题导致的导入失败问题

2.更改tools下数据库相关用户名以及密码配置情况

## 白箱测试

1. 在app中已经登陆有一个账户时，后端并未检测这个用户是否存在。（大概率是app问题）

2. 点赞功能正常
3. 喜欢功能正常

4. 关注：（新注册用户）在关注一个人后，在主页观察到人数变化，但是点进去无法获得具体人数

5. 评论功能正常 
   * 输入正常评论输入  ok
   * 输入emoji表情        ok
   * 评论删除功能正常，数据库操作正常

## 安全性测试

1.登录密码

登陆密码存入数据库时，明文保存（可改进为加密保存）。

##### ** （测试时发现问题1：可以以多个空格的方式作为密码来注册登录，并不符合常理，建议查看这里是否能够添加相关判断操作）**
##### ** （测试时发现问题2：可以以多个空格开头的方式以任意字符来结尾作为来登录登录，不符合常理，建议查看这里是否能够添加相关判断操作，例如判断限定为邮箱地址）**
数据库的密码在服务端写入，若为云端数据库，在开源时有数据库被爆破的风险。

2.sql注入测试

+ 经过测试，查看源码后暂时未发现有效的注入手段

其主要原因为，在前端能够操作的部分为两个方面：

1. 注册
2. 添加评论
##### ** （测试时发现问题3：可以以纯空格的方式来回复，建议查看这里是否能够添加相关判断操作，例如判断限定为不能以纯空格方式回复）**
##### **  (测试时发现可优化部分：同时若评论以空格为开头，将其空格去除，使得评论能够顶格显示并保存在数据库中，eg. ltrim*()) **
+ 而以上两个功能，并不提供查询功能，大部分客户端的键入都会由insert语句输入，且，查询环节并不需要通过我们来输入某一个参数来获得某一组数据，或是结果。

+ 但是仍然在可能的位置进行如下注入分析以及尝试：

1. + 在服务端最敏感的表 user 表返回语句处进行测试，采取最近点的 sql注入语句 ： 设置用户为 1' or '1'='1'  ，密码为 1
      这时，由于客户端在登陆时，后台服务器端不允许密码长度小于 6 ，无法通过sql注入创造一个万能用户

2. + 评论的返回与写入

   1. 评论的get方法，利用的为视频id，这个位置并不能由用户处理，故无法进行相关注入
   2. 

即 ：在测试的范围内未检查出有效的sql注入手段（有待继续测试）
       * 待测试问题 (是否存在可能性，将表中所有的用户名以及密码都进行窜改？？将密码全部篡改为空格，账号篡改为空格开头以sql语句来查询账户userId号
       并拼接在来结尾的账户名的手段来实现，在不确定userId的情况下获得来以一个特定的万能用户？？) *

		
