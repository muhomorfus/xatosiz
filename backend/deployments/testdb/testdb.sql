insert into "group" (uuid) values
('b4b17b07-ecd4-45ba-84cc-9708be9eba6f'), -- 1
('567af5eb-192a-4dc0-966a-d5755e248307'), -- 2
('f9d5aca1-becb-4812-95a4-659e8d1d9160'), -- 3
('8e5ba9ce-1d63-4bfb-92f4-15fa1b7f71b5'); -- 4

insert into trace (uuid, group_uuid, parent_uuid, title, time_start, time_end, component) values
('f2751424-3a3f-4515-bb3e-507eb2032110', 'b4b17b07-ecd4-45ba-84cc-9708be9eba6f', null, 'Title 1', current_timestamp - '10 minutes'::interval, null, 'service-1'), -- group = 1
('888eaedd-7855-4a9c-b6ab-45ba76123b01', 'b4b17b07-ecd4-45ba-84cc-9708be9eba6f', 'f2751424-3a3f-4515-bb3e-507eb2032110', 'Title 2', current_timestamp - '7 minutes'::interval, null, 'service-2'), -- group = 1
('9be1f78f-6df3-43ea-ab0c-f7c58c6e675b', 'b4b17b07-ecd4-45ba-84cc-9708be9eba6f', 'f2751424-3a3f-4515-bb3e-507eb2032110', 'Title 3', current_timestamp - '6 minutes'::interval, current_timestamp - '3 minutes'::interval, 'service-2'), -- group = 1
('eb0f3f3c-00ce-4f73-8e9a-67c784d08b15', 'b4b17b07-ecd4-45ba-84cc-9708be9eba6f', '9be1f78f-6df3-43ea-ab0c-f7c58c6e675b', 'Title 4', current_timestamp - '5 minutes'::interval, current_timestamp - '4 minutes'::interval, 'service-2'), -- group = 1
('fd2be14d-eeb7-45b7-84a5-bdbc16cd4ba5', '567af5eb-192a-4dc0-966a-d5755e248307', null, 'Title 5', current_timestamp - '10 minutes'::interval, null, 'service-3'), -- group = 2
('008a4c30-27eb-4f6d-bf99-8f6a7548b260', 'f9d5aca1-becb-4812-95a4-659e8d1d9160', null, 'Title 6', current_timestamp - '5 minutes'::interval, current_timestamp - '2 minutes'::interval, 'service-4'); -- group = 3

insert into event (uuid, trace_uuid, time, priority, message, payload, fixed) values
('a60089bb-6e41-4c43-9387-e12b87e72251', 'f2751424-3a3f-4515-bb3e-507eb2032110', current_timestamp - '6 minutes'::interval, 'info', 'Event 1', null, false), -- group = 1, trace = 1
('7e7c9d58-8e43-4d10-b077-09fce8854efc', 'f2751424-3a3f-4515-bb3e-507eb2032110', current_timestamp - '5 minutes'::interval, 'error', 'Event 2', null, false), -- group = 1, trace = 1
('fe672053-576c-4a76-9e5f-32e9c74eb5a0', '888eaedd-7855-4a9c-b6ab-45ba76123b01', current_timestamp - '5 minutes'::interval, 'warning', 'Event 3', null, true), -- group = 1, trace = 2
('cd8a1d49-e382-4d6e-ad21-39a271709dd6', '008a4c30-27eb-4f6d-bf99-8f6a7548b260', current_timestamp - '4 minutes'::interval, 'warning', 'Event 4', null, true), -- group = 3, trace = 6
('3bba9363-2411-4c4a-b830-536bdfcdbec0', 'fd2be14d-eeb7-45b7-84a5-bdbc16cd4ba5', current_timestamp - '8 minutes'::interval, 'warning', 'Event 5', null, false), -- group = 2, trace = 5
('e52f8e00-d41e-4536-b392-1fbd5e372ed5', 'fd2be14d-eeb7-45b7-84a5-bdbc16cd4ba5', current_timestamp - '7 minutes'::interval, 'info', 'Event 6', null, false), -- group = 2, trace = 5
('0e1520a3-cce7-43e6-bef3-6bb3a5cba11c', 'fd2be14d-eeb7-45b7-84a5-bdbc16cd4ba5', current_timestamp - '6 minutes'::interval, 'fatal', 'Event 7', null, true), -- group = 2, trace = 5
('fffd9eb9-af2b-4c82-a7ca-9f31274fa050', 'fd2be14d-eeb7-45b7-84a5-bdbc16cd4ba5', current_timestamp - '5 minutes'::interval, 'error', 'Event 8', null, false); -- group = 2, trace = 5