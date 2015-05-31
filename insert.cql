//test values
use alexandria;

insert into origins (originid, name) values (3c37dc44-908b-4f7f-b17b-6363904264f3, 'velocillama.com.');
insert into origins (originid, name) values (642f84f9-161a-4463-a622-86622bfba8d5, 'ronin.io.');

insert into rr (id, origin, type, name, value, ttl) values (134084c0-bfef-4bf3-9dcd-e9881e0e530a, 642f84f9-161a-4463-a622-86622bfba8d5, 'A', 'ronin.io.', '127.0.0.1', 600);
insert into rr (id, origin, type, name, value, ttl) values (ccc084c0-ccef-cbf3-cdcd-c9881e0e530c, 642f84f9-161a-4463-a622-86622bfba8d5, 'CNAME', 'www.ronin.io.', 'ronin.io.', 600);

insert into originrr (id, rr) values (3c37dc44-908b-4f7f-b17b-6363904264f3, 134084c0-bfef-4bf3-9dcd-e9881e0e530a);
insert into originrr (id, rr) values (3c37dc44-908b-4f7f-b17b-6363904264f3, ccc084c0-ccef-cbf3-cdcd-c9881e0e530c);

insert into qanswer (question, answer) values ('A-ronin.io.', 'ronin.io. 600 IN A 127.0.0.1');
insert into qanswer (question, answer) values ('CNAME-www.ronin.io.', 'www 600 IN CNAME ronin.io.');
insert into qanswer (question, answer) values ('A-multi.ronin.io.', 'multi.ronin.io. 600 IN A 127.0.0.1');
insert into qanswer (question, answer) values ('A-multi.ronin.io.', 'multi.ronin.io. 600 IN A 127.0.0.2');
