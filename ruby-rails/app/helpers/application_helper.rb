module ApplicationHelper
  # Converts active record date time to HTML5 datetime-local format. https://developer.mozilla.org/en-US/docs/Web/HTML/Element/input/datetime-local
  # Params:
  # +date+:: active record date time
  def convert_to_html_date_time_format(date)
    date.nil? ? nil : date.to_datetime.rfc3339[0..-7]
  end
end
