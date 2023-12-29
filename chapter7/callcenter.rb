class CallCenter
  def initialize(respondent_count, manager_count, director_count)
    @respondents = Array.new(respondent_count) { Respondent.new(self) }
    @managers = Array.new(manager_count) { Manager.new(self) }
    @directors = Array.new(director_count) { Director.new(self) }
    @respondents_queue = []
    @managers_queue = []
    @directors_queue = []
  end

  def assign_call(employee_type, call)
    avail, wait = get_queues(employee_type)

    if avail.empty?
      wait.push(call)
    else
      avail.pop.assign_call(call)
    end
  end

  def reassign(employee)
    avail, wait = get_queues(employee.type)

    if wait.empty?
      avail.push(employee)
    else
      employee.assign_call(wait.shift)
    end
  end

  def dispatch_call(call)
    assign_call(:respondent, call)
  end

  private

  def get_queues(employee_type)
    case employee_type
    when :respondent
      [@respondents, @respondents_queue]
    when :manager
      [@managers, @managers_queue]
    when :director
      [@directors, @directors_queue]
    else
      raise "Invalid employee type #{employee_type}"
    end
  end
end

class Employee
  def initialize(call_center)
    @call_center = call_center
  end

  def assign_call(call)
    raise "Class #{self.class} must implement #assign_call"
  end

  def type
    raise "Class #{self.class} must implement #type"
  end

  private

  def can_handle_call?
    rand(5) != 0
  end

  def escalate_call(call)
    raise "Class #{self.class} must implement #escalate_call"
  end

  def handle_call(call)
    raise "Class #{self.class} must implement #handle_call"
  end
end

class Respondent < Employee
  def assign_call(call)
    if can_handle_call?
      handle_call(call)
    else
      escalate_call(call)
    end
    @call_center.reassign(self)
  end

  def type
    :respondent
  end

  private

  def escalate_call(call)
    @call_center.assign_call(:manager, call)
  end

  def handle_call(call)
    puts "Respondent #{self} handled call #{call}"
  end
end

class Manager < Employee
  def assign_call(call)
    if can_handle_call?
      handle_call(call)
    else
      escalate_call(call)
    end
    @call_center.reassign(self)
  end

  def type
    :manager
  end

  private

  def escalate_call(call)
    @call_center.assign_call(:director, call)
  end

  def handle_call(call)
    puts "Manager #{self} handled call #{call}"
  end
end

class Director < Employee
  def assign_call(call)
    handle_call(call)
    @call_center.reassign(self)
  end

  def type
    :director
  end

  private

  def escalate_call(call)
    raise "Director #{self} cannot escalate call #{call}"
  end

  def handle_call(call)
    puts "Director #{self} handled call #{call}"
  end
end

cs = CallCenter.new(5, 2, 1)
20.times { |i| cs.dispatch_call("Call #{i}") }
