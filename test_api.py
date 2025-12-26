import requests
import json

BASE_URL = "http://localhost:8080/v1/api"

def test_blog_api():
    # 测试获取博客列表
    try:
        response = requests.get(f"{BASE_URL}/blog/")
        print(f"博客列表接口状态: {response.status_code}")
        if response.status_code == 200:
            print("博客列表接口正常")
            data = response.json()
            print(f"返回数据: {json.dumps(data, ensure_ascii=False, indent=2)}")
        else:
            print(f"错误: {response.text}")
    except Exception as e:
        print(f"博客列表接口测试失败: {e}")

if __name__ == "__main__":
    test_blog_api()