INSERT INTO "todo"."organizations" (id, name) VALUES
    ('7f0357c3-6719-4836-a3eb-bcf06ec2d759', 'A, Inc.');

INSERT INTO "todo"."roles" (id, name) VALUES
    ('11111111-1111-1111-1111-111111111111', 'none'),
    ('22222222-2222-2222-2222-222222222222', 'viewer'),
    ('33333333-3333-3333-3333-333333333333', 'editor');

INSERT INTO "todo"."users" (id, name, organization_id) VALUES
    ('fa224131-4ac9-4bc1-ae14-7d5f2c226255', 'Alice / Viewer', '7f0357c3-6719-4836-a3eb-bcf06ec2d759'),
    ('6d2dfe34-f2a5-4d34-b865-c2457062cec5', 'Bob / Editor', '7f0357c3-6719-4836-a3eb-bcf06ec2d759');

INSERT INTO "todo"."user_roles" (user_id, role_id) VALUES
    ('fa224131-4ac9-4bc1-ae14-7d5f2c226255', '22222222-2222-2222-2222-222222222222'), -- Alice: viewer
    ('6d2dfe34-f2a5-4d34-b865-c2457062cec5', '33333333-3333-3333-3333-333333333333'); -- Bob: editor
