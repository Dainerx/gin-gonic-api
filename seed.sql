INSERT INTO roles(id, "role", created_at, updated_at, deleted_at)
VALUES
(1, 'ADMIN', current_timestamp, null, null),
(2, 'USER', current_timestamp, null, null);

INSERT INTO users (name, email, password, role_id) VALUES
('John Doe', 'john.doe@example.com', 'hashedpassword1', 1),
('Jane Doe', 'jane.doe@example.com', 'hashedpassword2', 2),
('Bob Smith', 'bob.smith@example.com', 'hashedpassword3', 1),
('Alice Johnson', 'alice.johnson@example.com', 'hashedpassword4', 2);


-- Insert models
INSERT INTO models (name, description, thumbnail, private_score, community_score) VALUES
('Fibrosis Detection 1.0', 'Pathy internal model for fibrosis quanitfication in Kidney slides stained with Masson''s Trichrome', NULL, 0.89, 0.83);

-- Insert projects
INSERT INTO projects (name, description, thumbnail) VALUES
('ohio-state-project', 'Sanjil ohio state project', NULL),
('test-project', 'Pathy test project', NULL);

-- Get the ID of the 'ohio-state-project'
DO $$ DECLARE project_id integer;
BEGIN
   SELECT id INTO project_id FROM projects WHERE name = 'ohio-state-project';

   -- Insert slides
   INSERT INTO slides (file_path, thumbnail, metadata, project_id) VALUES
   ('first.svs', NULL, false, project_id),
   ('second.svs', NULL, false, project_id),
END $$;