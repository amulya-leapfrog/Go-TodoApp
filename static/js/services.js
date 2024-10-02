export async function fetchTaskById(id) {
  let payload = {
    task: "",
    status: "",
  };

  let taskId = 0;

  try {
    const response = await fetch(`/todo/${id}`);
    if (!response.ok) {
      throw new Error('Network response was not ok');
    }

    const data = await response.json();
    payload = {
      ...payload,
      task: data.task,
      status: data.status,
    };
    taskId = data.id;
  } catch (error) {
    alert("Something Went Wrong!");
    console.error("Error fetching todo:", error);
  }

  return { payload, taskId };
}

export async function createTodo(payload) {
  try {
    const response = await fetch(`/todo/`, {
      method: 'POST',
      headers: {
        'Content-Type': 'application/json',
      },
      body: JSON.stringify(payload),
    });

    if (!response.ok) {
      throw new Error('Network response was not ok');
    }

    window.location.reload();
  } catch (error) {
    alert("Something Went Wrong!");
    console.error("Error creating task:", error);
  }
}

export async function updateStatus(id, payload) {
  try {
    const response = await fetch(`/todo/${id}`, {
      method: 'PUT',
      headers: {
        'Content-Type': 'application/json',
      },
      body: JSON.stringify(payload),
    });

    if (!response.ok) {
      throw new Error('Network response was not ok');
    }

    window.location.reload();
  } catch (error) {
    alert("Something Went Wrong!");
    console.error("Error updating status:", error);
  }
}

export async function updateAllTodos() {
  try {
    const response = await fetch(`/todo/`, {
      method: 'PATCH',
    });

    if (!response.ok) {
      throw new Error('Network response was not ok');
    }

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
    const response = await fetch(`/todo/${taskId}`, {
      method: 'PUT',
      headers: {
        'Content-Type': 'application/json',
      },
      body: JSON.stringify(payload),
    });

    if (!response.ok) {
      throw new Error('Network response was not ok');
    }

    window.location.reload();
  } catch (error) {
    alert("Something Went Wrong!");
    console.error("Error updating task:", error);
  }
}

export async function deleteTodo(id) {
  try {
    const response = await fetch(`/todo/${id}`, {
      method: 'DELETE',
    });

    if (!response.ok) {
      throw new Error('Network response was not ok');
    }

    window.location.reload();
  } catch (error) {
    alert("Something Went Wrong!");
    console.error("Error deleting task:", error);
  }
}

export async function deleteAllTodos() {
  try {
    const response = await fetch(`/todo/`, {
      method: 'DELETE',
    });

    if (!response.ok) {
      throw new Error('Network response was not ok');
    }

    window.location.reload();
  } catch (error) {
    alert("Something Went Wrong!");
    console.error("Error deleting tasks:", error);
  }
}
