CREATE OR REPLACE FUNCTION validate_sub_dates(start_d DATE, end_d DATE)
RETURNS BOOLEAN AS $$
BEGIN
  IF end_d IS NULL THEN RETURN TRUE; END IF;
  IF end_d >= start_d + INTERVAL '1 month' THEN RETURN TRUE; END IF;

  RAISE EXCEPTION USING
    ERRCODE = '23514',
    MESSAGE = 'end date should be later than start date for at least a month';
END;
$$ LANGUAGE plpgsql;

ALTER TABLE subs DROP CONSTRAINT IF EXISTS check_date_range;

ALTER TABLE subs
ADD CONSTRAINT check_date_range CHECK (
    validate_sub_dates (start_date, end_date)
);