### Go学习
<br>     

#### 参考资料
1. [Go 语言上手 - 基础语言](https://juejin.cn/post/7093721879462019102)
2. [Go语言圣经](https://books.studygolang.com/gopl-zh/)
3. [青训营完整手册](https://bytedance.feishu.cn/docs/doccnFRB1TXYJPK6yprPETHLXgd#q8ZYps)
4. [抖音项目-青训营](https://bytedance.feishu.cn/docx/doxcnbgkMy2J0Y3E6ihqrvtHXPg)
5. ...

<br>

#### 实践
1. 视频Feed流、视频投稿、个人信息：    

&emsp;&emsp;支持所有用户刷抖音，按投稿时间倒叙推出，登录用户可自己拍视频投稿，查看自己的基本信息和投稿列表，注册用户流程简化。
   - 支持所有用户刷抖音，按投稿时间倒叙推出，单次最多30个视频
   - 注册登录功能  
   - 实现登录用户投稿功能
   - 查看自己基本信息和投稿列表

根据分析，第一步应该设计数据库的表结构：
   - 首先实现基础功能
   - 应包含用户表、视频表

**第一部分已全部完成**   
*存在问题： 1.只有重启程序才能刷新首页，没有刷新功能（新版apk已解决该问题） 2.封面图没有上传* 
   
1. 点赞列表，用户评论
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

**点赞功能已完成：**
- 包括视频点赞显示正确点赞数，点赞加一，取消赞减一，点赞实时修改数据库
- 以及用户喜爱列表显示正确的喜爱视频
- 用户登录时正确显示红心状态
- 喜欢后显示正确的数量
- 总获赞显示正确的数量

问题：
- 登录之后没有自动刷新feed流，需要手动刷新后，才能正确显示红心状态
- 总获赞数只有登录时才能刷新
- 作品数始终为零（修改：添加字段work_count） **完成字段添加，增加作品数变化，代码发生大量修改，体现预留字段的重要性**
- 新的问题，作品数似乎无法像喜欢数那样修改后立即变化，只能登录时变化一次
- 个人信息页面作品数和总获赞数似乎无法即时刷新，而关注，粉丝，喜欢可以即时刷新

**评论功能已完成：**
- 数据库内容的补充
- 包括不同视频显示对应评论内容
- 评论发出以及评论删除

问题：
- 评论完成后没有自动调用评论接口，导致需要关闭评论页面重新打开才能刷新 **(该问题已解决，评论时需要在返回值中传递评论内容)**


1. 关注列表，粉丝列表
- 关注功能
  - 和点赞功能类似， 点击+关注时，将视频的发布用户加入登录用户的关注表中，没有关注表则动态建立关注表，同时修改用户关注数，和对应用户粉丝数，将登录用户加入到发布视频用户的粉丝表中，没有粉丝表则动态生成粉丝表，修改发布用户粉丝数
- 关注数和粉丝数加入用户信息中
  - 数据库的user表需要添加关注数和粉丝数
- 能够显示关注列表和粉丝列表
  - 显示关注列表和粉丝列表只需要读数据库对应表就行
  需要注意的是初始化视频显示时，需判断“+”即关注符号需不需要显示，和点赞功能类似，读取登录用户的关注列表判断是否应该显示“+”
  - 表名规则同上

**关注以及粉丝功能已完成：**
- 完成数据表的扩充
- 登录时个人信息页面显示
- 关注时即时修改数据表内容
- 所有页面数据同步

问题：
- 和点赞类似，登录后需刷新才能看到关注状态。。。
- 因为登录后需要手动刷新才能显示点赞以及关注状态，所以在未刷新状态下，可以对已关注用户，已点赞视频再次点赞再次关注  
 （后端加入点赞，关注前判断是否重复操作，若重复直接提示 "repeat operation"，并且不执行对应操作，最好的解决方式是前端增加一个刷新接口，需要时调用，刷新视频数据）  
 *不能通过直接改变数据库属性来解决这个问题，虽然禁止重复可以防止插入重复数据，但是没办法解决其它表中计数加1的问题*
 **已完成后端点赞时，重复操作判断**

<br>

  &emsp;&emsp;初始数据库文件在sql文件夹中，运行程序前需先配置好数据库环境，将数据导入数据库中 [数据库](sql/douyin.sql)

  **提示：连接数据库时记得修改代码中的数据库账户密码，在tool文件中**
