-- Initialize mock data for the services table
INSERT INTO public.services (id, email, ip_address)
VALUES 
    ('c1a72b8c-8f6e-401d-94ef-3f1d527760cc', 'service1@example.com', '192.168.1.1'),
    ('d1a53c9d-9a2f-482e-bd8e-4e2d64c018a2', 'service2@example.com', '192.168.1.2');

-- Initialize mock data for the queues table
INSERT INTO public.queues (id, position, service_id)
VALUES 
    ('fa5f9fbc-3d8d-4b56-8f1e-3c6a4a37b2a7', 0, 'c1a72b8c-8f6e-401d-94ef-3f1d527760cc'),
    ('bb8390d9-c9d4-44a2-bd1e-57c1369ef6a3', 1, 'c1a72b8c-8f6e-401d-94ef-3f1d527760cc'),
    ('ac529193-4978-4db6-bd59-6e98366fae3e', 2, 'c1a72b8c-8f6e-401d-94ef-3f1d527760cc'),
    ('a1221a99-f1c4-46a8-8d1b-d3d94745dc2d', 0, 'd1a53c9d-9a2f-482e-bd8e-4e2d64c018a2'),
    ('b1122a88-01a7-4fa8-9f8e-e3b57245ad8c', 1, 'd1a53c9d-9a2f-482e-bd8e-4e2d64c018a2'),
    ('c2233b77-12b3-4fb8-9f9f-f3c58245bc9d', 2, 'd1a53c9d-9a2f-482e-bd8e-4e2d64c018a2'),
    ('d3344c66-23b6-4fa9-8f9f-04c69245cdad', 3, 'c1a72b8c-8f6e-401d-94ef-3f1d527760cc'),
    ('e4455d55-34b9-4faa-8faf-15d7a245debe', 3, 'd1a53c9d-9a2f-482e-bd8e-4e2d64c018a2'),
    ('f5566e44-45bc-4fab-9fbf-26d8b245efcf', 4, 'c1a72b8c-8f6e-401d-94ef-3f1d527760cc'),
    ('f6677f33-56bf-4fac-afaf-37d9c2450fe0', 4, 'd1a53c9d-9a2f-482e-bd8e-4e2d64c018a2');
