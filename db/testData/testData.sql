-- SQL script to insert clients into the clients table

-- user
INSERT INTO users (
    id,
    fname,
    lname,
    email,
    "hash",
    nameforheader,
    license,
    created_at,
    updated_at
) VALUES (
    '1b486bf2-a886-416e-99c7-b3be602a9b3e',
    'John',
    'Doe',
    'jdoe123@gmail.com',
    -- password is 'password12345'
    '$2a$10$Bs2i7sVC4EhHdGEAFMM8UuCsgdbeSeo2qruEZ9hQ6uKj3fprHu4ii',
    'J. Doe',
    'MF 416',
    '2024-10-21 12:04:22.120231',
    '2024-10-21 12:04:30.093619'
) ON CONFLICT (id) DO NOTHING;
-- clients
INSERT INTO clients (
    id, 
    user_id, 
    fname, 
    lname, 
    email, 
    phone, 
    balance, 
    balancenotifythreshold, 
    rate, 
    isarchived, 
    created_at, 
    updated_at
) VALUES
('f720d216-4a2e-42f8-a95b-4d6678d397e1', '1b486bf2-a886-416e-99c7-b3be602a9b3e', 'Mark', 'Twain', 'mtwainthman@gmail.com', '222-222-2222', 0, 10000, 25000, false, '2024-11-22 22:38:58.107718', '2024-11-22 22:38:58.107718'),
('f720d216-4a2e-42f8-a95b-4d6678d397e2', '1b486bf2-a886-416e-99c7-b3be602a9b3e', 'John', 'Doe', 'jdoe123@gmail.com', '222-222-2222', 250000, 100000, 25000, false, '2024-10-29 22:38:58.107718', '2024-10-29 22:38:58.107718'),
('f720d216-4a4e-42f8-a92b-4d6978d397e3', '1b486bf2-a886-416e-99c7-b3be602a9b3e', 'Jane', 'Doe', 'jane.doe@example.com', '333-333-3333', 200000, 50000, 18000, false, '2024-10-30 15:45:10.107718', '2024-10-30 15:45:10.107718'),
('d720d216-7a2e-42f8-c95b-7d5678d397e1', '1b486bf2-a886-416e-99c7-b3be602a9b3e', 'Charles', 'Dickens', 'charles.d@example.com', '555-555-5555', 0, 100000, 23000, false, '2024-11-03 08:30:58.107718', '2024-11-03 08:30:58.107718'),
('b345d216-5a2e-42f8-b95b-5d5678d297e1', '1b486bf2-a886-416e-99c7-b3be602a9b3e', 'Jane', 'Austen', 'jane.austen@example.com', '333-333-3333', 120000, 50000, 15000, false, '2024-11-01 15:20:15.107718', '2024-11-01 15:20:15.107718'),
('c920d216-6b2e-42f8-b95b-5d6678d297e1', '1b486bf2-a886-416e-99c7-b3be602a9b3e', 'Emily', 'Bronte', 'emily.bronte@example.com', '444-444-4444', 200000, 80000, 18000, false, '2024-11-02 10:10:10.107718', '2024-11-02 10:10:10.107718'),
('e820d216-8b2e-42f8-d95b-8d5678d397e1', '1b486bf2-a886-416e-99c7-b3be602a9b3e', 'William', 'Shakespeare', 'william.s@example.com', '666-666-6666', 500000, 200000, 30000, false, '2024-11-04 12:00:58.107718', '2024-11-04 12:00:58.107718'),
('f810d236-1c4a-47fb-998e-a7df77b29856', '1b486bf2-a886-416e-99c7-b3be602a9b3e', 'Leo', 'Tolstoy', 'leo.t@example.com', '777-777-7777', 150000, 60000, 22000, false, '2024-11-05 14:05:25.107718', '2024-11-05 14:05:25.107718'),
('fb96b3d7-38e6-4bc5-a663-8bc0ae1488a7', '1b486bf2-a886-416e-99c7-b3be602a9b3e', 'Virginia', 'Woolf', 'virginia.woolf@example.com', '888-888-8888', 180000, 70000, 24000, false, '2024-11-06 10:15:30.107718', '2024-11-06 10:15:30.107718'),
('c8f2337e-d023-435f-b232-16d71b2cfe9c', '1b486bf2-a886-416e-99c7-b3be602a9b3e', 'Herman', 'Melville', 'herman.m@example.com', '999-999-9999', 300000, 100000, 26000, false, '2024-11-07 09:20:00.107718', '2024-11-07 09:20:00.107718'),
('ba56f3e4-d19b-43b5-9823-87b7ab08fa79', '1b486bf2-a886-416e-99c7-b3be602a9b3e', 'George', 'Eliot', 'george.eliot@example.com', '101-101-1010', 400000, 150000, 27000, false, '2024-11-08 13:45:50.107718', '2024-11-08 13:45:50.107718'),
('a352d5ec-2a24-411b-91a8-e75a9625b037', '1b486bf2-a886-416e-99c7-b3be602a9b3e', 'Toni', 'Morrison', 'toni.m@example.com', '202-202-2022', 220000, 90000, 29000, false, '2024-11-09 16:30:45.107718', '2024-11-09 16:30:45.107718'),
('f842d56e-3e67-4538-85a9-40b1d56c7e0b', '1b486bf2-a886-416e-99c7-b3be602a9b3e', 'Fyodor', 'Dostoevsky', 'fyodor.d@example.com', '303-303-3033', 170000, 80000, 25000, false, '2024-11-10 11:55:20.107718', '2024-11-10 11:55:20.107718'),
('ac43d2f6-6e14-4ec5-9e9c-32a73249e44a', '1b486bf2-a886-416e-99c7-b3be602a9b3e', 'Gabriel', 'Garcia Marquez', 'gabriel.g@example.com', '404-404-4044', 250000, 100000, 28000, false, '2024-11-11 17:25:35.107718', '2024-11-11 17:25:35.107718'),
('de56f73a-8a39-4e9b-bfd5-caf9b7a2e013', '1b486bf2-a886-416e-99c7-b3be602a9b3e', 'James', 'Joyce', 'james.j@example.com', '505-505-5055', 130000, 40000, 21000, false, '2024-11-12 14:40:50.107718', '2024-11-12 14:40:50.107718')
ON CONFLICT (id) DO NOTHING;


