export async function fetchTaskById(id) {
  let payload = {
    task: "",
    status: "",
  };

  let taskId = 0;

  try {
    const response = await axios.get(`http://localhost:8000/todo/${id}`);
    payload = {
      ...payload,
      task: response.data.task,
      status: response.data.status,
    };
    taskId = response.data.id;
  } catch (error) {
    alert("Something Went Wrong!");
    console.error("Error fetching todo:", error);
  }

  return { payload, taskId };
}

export async function createTodo(payload) {
  try {
    await axios.post(`http://localhost:8000/todo/`, payload);
    window.location.reload();
  } catch (error) {
    alert("Something Went Wrong!");
    console.error("Error creating task:", error);
  }
}

export async function updateStatus(id, payload) {
  try {
    await axios.put(`http://localhost:8000/todo/${id}`, payload);
    window.location.reload();
  } catch (error) {
    alert("Something Went Wrong!");
    console.error("Error updating status:", error);
  }
}

export async function updateAllTodos() {
  try {
    await axios.patch("http://localhost:8000/todo/");
    window.location.reload();
  } catch (error) {
    alert("Something Went Wrong!");
    console.error("Error updating tasks:", error);
  }
}

export async function handleEdit(event) {
  event.preventDefault();

  const form = document.querySelector("form");
  const taskId = form.getAttribute("data-task-id");
  const taskName = document.getElementById("editTask").value;
  const status = document.querySelector('input[name="status"]:checked').value;

  const payload = {
    task: taskName,
    status: status,
  };

  try {
    await axios.put(`http://localhost:8000/todo/${taskId}`, payload);
    window.location.reload();
  } catch (error) {
    alert("Something Went Wrong!");
    console.error("Error updating task:", error);
  }
}

export async function deleteTodo(id) {
  try {
    await axios.delete(`http://localhost:8000/todo/${id}`);
    window.location.reload();
  } catch (error) {
    alert("Something Went Wrong!");
    console.error("Error deleting task:", error);
  }
}

export async function deleteAllTodos() {
  try {
    await axios.delete("http://localhost:8000/todo/");
    window.location.reload();
  } catch (error) {
    alert("Something Went Wrong!");
    console.error("Error deleting tasks:", error);
  }
}
