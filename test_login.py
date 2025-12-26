import requests
import json

BASE_URL = "http://localhost:8080/v1/api"

def test_login_api():
    # 测试登录接口
    try:
        payload = {
            "username": "testuser",
            "password": "testpassword"
        }
        
        response = requests.post(f"{BASE_URL}/user/login", json=payload)
        print(f"登录接口状态: {response.status_code}")
        print(f"响应头: {dict(response.headers)}")
        
        if response.status_code == 200:
            data = response.json()
            print("登录接口正常")
            print(f"返回数据: {json.dumps(data, ensure_ascii=False, indent=2)}")
        else:
            print(f"错误响应: {response.text}")
            print(f"错误详情: {response.reason}")
            
    except Exception as e:
        print(f"登录接口测试失败: {e}")

if __name__ == "__main__":
    test_login_api()