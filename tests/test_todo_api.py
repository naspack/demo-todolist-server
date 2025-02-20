import unittest
import requests
import time
from typing import Dict

class TestTodoAPI(unittest.TestCase):
    BASE_URL = "http://localhost:8080"

    def setUp(self):
        self.todo_payload = {
            "title": "Test Todo",
            "description": "This is a test todo item"
        }
        # Create a todo for tests that need an existing todo
        response = requests.post(f"{self.BASE_URL}/api/todos", json=self.todo_payload)
        self.assertEqual(response.status_code, 201)
        self.created_todo = response.json()

    def test_create_todo(self):
        """Test creating a new todo"""
        response = requests.post(f"{self.BASE_URL}/api/todos", json=self.todo_payload)
        self.assertEqual(response.status_code, 201)
        data = response.json()
        
        self.assertEqual(data["title"], self.todo_payload["title"])
        self.assertEqual(data["description"], self.todo_payload["description"])
        self.assertFalse(data["completed"])
        self.assertIn("ID", data)
        self.assertIn("CreatedAt", data)
        self.assertIn("UpdatedAt", data)

    def test_get_todos(self):
        """Test getting all todos"""
        response = requests.get(f"{self.BASE_URL}/api/todos")
        self.assertEqual(response.status_code, 200)
        todos = response.json()
        
        self.assertIsInstance(todos, list)
        self.assertTrue(any(todo["ID"] == self.created_todo["ID"] for todo in todos))

    def test_get_single_todo(self):
        """Test getting a single todo"""
        response = requests.get(f"{self.BASE_URL}/api/todos/{self.created_todo['ID']}")
        self.assertEqual(response.status_code, 200)
        todo = response.json()
        
        self.assertEqual(todo["ID"], self.created_todo["ID"])
        self.assertEqual(todo["title"], self.created_todo["title"])
        self.assertEqual(todo["description"], self.created_todo["description"])

    def test_update_todo(self):
        """Test updating a todo"""
        update_payload = {
            "title": "Updated Todo",
            "description": "This todo has been updated",
            "completed": True
        }
        
        response = requests.put(
            f"{self.BASE_URL}/api/todos/{self.created_todo['ID']}", 
            json=update_payload
        )
        self.assertEqual(response.status_code, 200)
        updated_todo = response.json()
        
        self.assertEqual(updated_todo["ID"], self.created_todo["ID"])
        self.assertEqual(updated_todo["title"], update_payload["title"])
        self.assertEqual(updated_todo["description"], update_payload["description"])
        self.assertTrue(updated_todo["completed"])

    def test_delete_todo(self):
        """Test deleting a todo"""
        response = requests.delete(f"{self.BASE_URL}/api/todos/{self.created_todo['ID']}")
        self.assertEqual(response.status_code, 204)
        
        # Verify the todo is deleted
        get_response = requests.get(f"{self.BASE_URL}/api/todos/{self.created_todo['ID']}")
        self.assertEqual(get_response.status_code, 404)

    def test_get_nonexistent_todo(self):
        """Test getting a todo that doesn't exist"""
        response = requests.get(f"{self.BASE_URL}/api/todos/99999")
        self.assertEqual(response.status_code, 404)
        self.assertIn("error", response.json())

    def test_create_todo_invalid_payload(self):
        """Test creating a todo with invalid payload"""
        response = requests.post(f"{self.BASE_URL}/api/todos", json={})
        self.assertEqual(response.status_code, 400)
        self.assertIn("error", response.json())

if __name__ == '__main__':
    unittest.main() 