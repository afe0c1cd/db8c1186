INSERT INTO todos (id, title, description, status, visibility, created_by_user_id, organization_id) VALUES
    ('11111111-1111-1111-1111-111111111111', '週次ミーティングの準備', '議題の整理と資料の準備', 'open', 'internal', '6d2dfe34-f2a5-4d34-b865-c2457062cec5', '7f0357c3-6719-4836-a3eb-bcf06ec2d759'),
    ('22222222-2222-2222-2222-222222222222', 'プロジェクト計画書の作成', 'Q3のプロジェクト計画をまとめる', 'open', 'internal', '6d2dfe34-f2a5-4d34-b865-c2457062cec5', '7f0357c3-6719-4836-a3eb-bcf06ec2d759'),
    ('33333333-3333-3333-3333-333333333333', 'コードレビューの実施', 'AliceのPRをレビューする', 'open', 'private', '6d2dfe34-f2a5-4d34-b865-c2457062cec5', '7f0357c3-6719-4836-a3eb-bcf06ec2d759'),
    ('44444444-4444-4444-4444-444444444444', '新機能の実装', 'ユーザー管理機能の追加', 'open', 'internal', 'fa224131-4ac9-4bc1-ae14-7d5f2c226255', '7f0357c3-6719-4836-a3eb-bcf06ec2d759');

INSERT INTO todo_assignees (user_id, todo_id, created_by_user_id) VALUES
    ('fa224131-4ac9-4bc1-ae14-7d5f2c226255', '11111111-1111-1111-1111-111111111111', '6d2dfe34-f2a5-4d34-b865-c2457062cec5'),
    ('6d2dfe34-f2a5-4d34-b865-c2457062cec5', '44444444-4444-4444-4444-444444444444', 'fa224131-4ac9-4bc1-ae14-7d5f2c226255');
