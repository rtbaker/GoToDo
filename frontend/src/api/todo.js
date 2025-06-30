const todoAPI=import.meta.env.VITE_API_URL + "/api/1.0/todos"

export async function deleteTodo(todo, successFn, errorFn, loginRequiredFn) {
    try {
            const response = await fetch(todoAPI + '/' + todo.id, {
              method: "DELETE",
              credentials: "include",
            });
        
            if (!response.ok) {
              console.log(`Response status: ${response.status}`);
    
              if (response.status === 401) {
                loginRequiredFn();
              }
    
              const json = await response.json();
              console.log(errorFn);
              console.log(json);
    
              errorFn(json.message);

              return;
            }
        
            // all good then close
            // Delete endpoint doesn't return any data so just fake some for consitency
            successFn({});
          } catch (error) {
            errorFn(error.message);
          }
}

export async function markTodoDone(todo, successFn, errorFn, loginRequiredFn) {
    try {
            const response = await fetch(todoAPI + '/' + todo.id, {
                body: JSON.stringify({ completed: true }),
                method: "PATCH",
                credentials: "include",
            });
        
            if (!response.ok) {
              console.log(`Response status: ${response.status}`);
    
              if (response.status === 401) {
                loginRequiredFn();
              }
    
              const json = await response.json();
              console.log(errorFn);
              console.log(json);
    
              errorFn(json.message);

              return;
            }
        
            // all good then close
            const json = await response.json();
            successFn(json);
          } catch (error) {
            errorFn(error.message);
          }
}