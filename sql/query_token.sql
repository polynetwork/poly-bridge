use polyswap;

set @tname='cat1';
set @tokens=('0c3c33da088abeee376418d3e384528c5aadba11', 'a85c9fc8f2c9060d674e0ca97f703a0a30619305', '455b51d882571e244d03668f1a458ca74e70d196');
select * from token_basics where name=@tname;
select * from tokens where token_basic_name=@tname;
select * from token_maps where 
src_token_hash in ('0c3c33da088abeee376418d3e384528c5aadba11', 'a85c9fc8f2c9060d674e0ca97f703a0a30619305', '455b51d882571e244d03668f1a458ca74e70d196')
and dst_token_hash in ('0c3c33da088abeee376418d3e384528c5aadba11', 'a85c9fc8f2c9060d674e0ca97f703a0a30619305', '455b51d882571e244d03668f1a458ca74e70d196')
;