CREATE OR REPLACE FUNCTION get_dogs()
RETURNS TABLE(
    id UUID, name VARCHAR,
    created_at TIMESTAMP WITH TIME ZONE,
    updated_at TIMESTAMP WITH TIME ZONE
) AS $$
BEGIN
    RETURN QUERY
    SELECT *
    FROM dogs;
END;
$$ LANGUAGE plpgsql;
