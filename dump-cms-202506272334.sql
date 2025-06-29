PGDMP      "                }            cms %   14.12 (Ubuntu 14.12-0ubuntu0.22.04.1)    17.2     !           0    0    ENCODING    ENCODING        SET client_encoding = 'UTF8';
                           false            "           0    0 
   STDSTRINGS 
   STDSTRINGS     (   SET standard_conforming_strings = 'on';
                           false            #           0    0 
   SEARCHPATH 
   SEARCHPATH     8   SELECT pg_catalog.set_config('search_path', '', false);
                           false            $           1262    16384    cms    DATABASE     n   CREATE DATABASE cms WITH TEMPLATE = template0 ENCODING = 'UTF8' LOCALE_PROVIDER = libc LOCALE = 'en_US.utf8';
    DROP DATABASE cms;
                     postgres    false                        2615    2200    public    SCHEMA        CREATE SCHEMA public;
    DROP SCHEMA public;
                     postgres    false            %           0    0    SCHEMA public    COMMENT     6   COMMENT ON SCHEMA public IS 'standard public schema';
                        postgres    false    5            &           0    0    SCHEMA public    ACL     Q   REVOKE USAGE ON SCHEMA public FROM PUBLIC;
GRANT ALL ON SCHEMA public TO PUBLIC;
                        postgres    false    5            �            1259    16385    article    TABLE     �   CREATE TABLE public.article (
    id integer NOT NULL,
    title character varying,
    content character varying,
    author_id integer,
    created_at timestamp without time zone,
    updated_at timestamp without time zone,
    tag jsonb
);
    DROP TABLE public.article;
       public         heap r       postgres    false    5            �            1259    16388    article_id_seq    SEQUENCE     �   CREATE SEQUENCE public.article_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;
 %   DROP SEQUENCE public.article_id_seq;
       public               postgres    false    210    5            '           0    0    article_id_seq    SEQUENCE OWNED BY     A   ALTER SEQUENCE public.article_id_seq OWNED BY public.article.id;
          public               postgres    false    211            �            1259    16430    payment    TABLE     <  CREATE TABLE public.payment (
    id uuid DEFAULT public.uuid_generate_v4() NOT NULL,
    amount numeric(10,2) NOT NULL,
    remarks character varying(255),
    balance_before numeric(10,2) NOT NULL,
    balance_after numeric(10,2) NOT NULL,
    created_date timestamp without time zone DEFAULT CURRENT_TIMESTAMP
);
    DROP TABLE public.payment;
       public         heap r       postgres    false    5    5    5            �            1259    16409    role    TABLE        CREATE TABLE public.role (
);
    DROP TABLE public.role;
       public         heap r       postgres    false    5            �            1259    16423    topup    TABLE       CREATE TABLE public.topup (
    id uuid DEFAULT public.uuid_generate_v4() NOT NULL,
    amount numeric(10,2) NOT NULL,
    balance_before numeric(10,2) NOT NULL,
    balance_after numeric(10,2) NOT NULL,
    created_date timestamp without time zone DEFAULT CURRENT_TIMESTAMP
);
    DROP TABLE public.topup;
       public         heap r       postgres    false    5    5    5            �            1259    16395    user    TABLE     �   CREATE TABLE public."user" (
    id integer NOT NULL,
    password character varying,
    email character varying,
    username character varying,
    created_at timestamp without time zone,
    updated_at timestamp without time zone
);
    DROP TABLE public."user";
       public         heap r       postgres    false    5            �            1259    16400    user_id_seq    SEQUENCE     �   CREATE SEQUENCE public.user_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;
 "   DROP SEQUENCE public.user_id_seq;
       public               postgres    false    5    212            (           0    0    user_id_seq    SEQUENCE OWNED BY     =   ALTER SEQUENCE public.user_id_seq OWNED BY public."user".id;
          public               postgres    false    213                       2604    16389 
   article id    DEFAULT     h   ALTER TABLE ONLY public.article ALTER COLUMN id SET DEFAULT nextval('public.article_id_seq'::regclass);
 9   ALTER TABLE public.article ALTER COLUMN id DROP DEFAULT;
       public               postgres    false    211    210            �           2604    16401    user id    DEFAULT     d   ALTER TABLE ONLY public."user" ALTER COLUMN id SET DEFAULT nextval('public.user_id_seq'::regclass);
 8   ALTER TABLE public."user" ALTER COLUMN id DROP DEFAULT;
       public               postgres    false    213    212                      0    16385    article 
   TABLE DATA           ]   COPY public.article (id, title, content, author_id, created_at, updated_at, tag) FROM stdin;
    public               postgres    false    210   r                 0    16430    payment 
   TABLE DATA           c   COPY public.payment (id, amount, remarks, balance_before, balance_after, created_date) FROM stdin;
    public               postgres    false    216   �                 0    16409    role 
   TABLE DATA              COPY public.role  FROM stdin;
    public               postgres    false    214   Q                 0    16423    topup 
   TABLE DATA           X   COPY public.topup (id, amount, balance_before, balance_after, created_date) FROM stdin;
    public               postgres    false    215   n                 0    16395    user 
   TABLE DATA           W   COPY public."user" (id, password, email, username, created_at, updated_at) FROM stdin;
    public               postgres    false    212   �       )           0    0    article_id_seq    SEQUENCE SET     <   SELECT pg_catalog.setval('public.article_id_seq', 1, true);
          public               postgres    false    211            *           0    0    user_id_seq    SEQUENCE SET     9   SELECT pg_catalog.setval('public.user_id_seq', 1, true);
          public               postgres    false    213            �           2606    16399    article article_pk 
   CONSTRAINT     P   ALTER TABLE ONLY public.article
    ADD CONSTRAINT article_pk PRIMARY KEY (id);
 <   ALTER TABLE ONLY public.article DROP CONSTRAINT article_pk;
       public                 postgres    false    210            �           2606    16436    payment payment_pkey 
   CONSTRAINT     R   ALTER TABLE ONLY public.payment
    ADD CONSTRAINT payment_pkey PRIMARY KEY (id);
 >   ALTER TABLE ONLY public.payment DROP CONSTRAINT payment_pkey;
       public                 postgres    false    216            �           2606    16429    topup topup_pkey 
   CONSTRAINT     N   ALTER TABLE ONLY public.topup
    ADD CONSTRAINT topup_pkey PRIMARY KEY (id);
 :   ALTER TABLE ONLY public.topup DROP CONSTRAINT topup_pkey;
       public                 postgres    false    215            �           2606    16408    user user_pk 
   CONSTRAINT     L   ALTER TABLE ONLY public."user"
    ADD CONSTRAINT user_pk PRIMARY KEY (id);
 8   ALTER TABLE ONLY public."user" DROP CONSTRAINT user_pk;
       public                 postgres    false    212               U   x�3�L��,,�,PH.�)H���/.)*MQH���UH���/JUHM�4�4202�50�52W04�26�24�365276�'������ 3�         j   x�%ɱ1�:�"8�q��)�&g����#A�����F3�mN��NE�����Z�������sO������
~0X	qK\W���ڋ�K�1~s]            x������ � �         k   x�����0�js�,@��גgI#*��#ėL�@��.�d��ڋ+e����]pė�B����C�r\�-�B�W�.p�L���˫�qbx������L:=} ĳ$Q         |   x�3�T1JT14P�.4v5v��)5�5,��2�6MҫJ*M	231p/���Ψ�7K+�ͶH���,�M��L/�K,O�sH�M���K�υ�p�����+[[XZ�[������ �U'�     