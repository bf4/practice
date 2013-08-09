class Bob

  # @param drivel [String] The drivel Bob is responding to
  # @return response [String] Bob's response to the drivel
  def hey(drivel)
    @drivel = drivel
    case
    when silence?   then annoyed
    when shouting?  then alarmed
    when asking?    then resigned
    else                 indifferent
    end
  end

  private

  def drivel
    @drivel
  end

  def silence?
    drivel.nil? || drivel.strip == ''
  end

  def annoyed
    "Fine. Be that way!"
  end

  def shouting?
    drivel.upcase == drivel
  end

  def alarmed
    "Woah, chill out!"
  end

  def asking?
    drivel.end_with?('?')
  end

  def resigned
    "Sure."
  end

  def indifferent
    "Whatever."
  end

end
