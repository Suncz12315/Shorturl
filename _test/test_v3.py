from locust import HttpLocust, TaskSet
import time

#根据长网址查询短网址
def queryltos(l):

    l.client.post("/query",{"longurl":"http://www.baidu.com"})

#根据短网址查询长网址
def querystol(l):

    l.client.post("/query",{"shorturl":"1"})

#短网址重定向
def redir(l):

    l.client.get("/redir/1")

#新的长网址映射短网址
def index(l):

    l.client.post("/query",{"longurl":"http://123.sogou.com"})

#def profile(l):

#    l.client.get("/test/?url=http://www.ckck.vip")

class UserBehavior(TaskSet):

    tasks = {redir:1,queryltos:1,querystol:1,index:1}

    def on_start(self):

        queryltos(self)

        querystol(self)

        index(self)


class WebsiteUser(HttpLocust):

    task_set = UserBehavior

    min_wait=1000

    max_wait=5000

    #host = "http://127.0.0.1:8080"
