use polyswap;

set @tname='SCAPES';
select * from token_basics where name=@tname;
select * from tokens where name=@tname;
select * from token_maps 
where 
src_token_hash in (select hash from tokens where name=@tname) or
dst_token_hash in (select hash from tokens where name=@tname);
