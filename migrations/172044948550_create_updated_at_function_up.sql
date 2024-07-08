CREATE OR REPLACE FUNCTION update_updated_at_column()
RETURNS TRIGGER as $$
BEGIN
    NEW.updated_at = now();
    RETURN NEW;
END;
$$ LANGUAGE plpgsql;
