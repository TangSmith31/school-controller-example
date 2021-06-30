# school-controller-example
Self-Define Controller Example: School & Student

## 背景描述
```
模拟一个新生登记系统，要求如下：
|- 登记一名新生的同时，该校的学生统计人数需要+1
|- 登记成功后，方可为该生下发学生证
|- 新生需持有学生证方可入校
```

## 安装
```bash
[root@centos1 ~]# git clone https://github.com/TangSmith31/school-controller-example
[root@centos1 ~]# cd school-controller-example/
[root@centos1 school-controller-example]# make
[root@centos1 school-controller-example]# make install
```

## 运行 & 测试
```bash
[root@centos1 school-controller-example]# kubectl apply -f config/samples/school.yaml 
[root@centos1 school-controller-example]# make run
```
模拟导入学生A、B、C
```bash
[root@centos1 school-controller-example]# kubectl apply -f config/samples/student_a.yaml
[root@centos1 school-controller-example]# kubectl apply -f config/samples/student_b.yaml
[root@centos1 school-controller-example]# kubectl apply -f config/samples/student_c.yaml
```