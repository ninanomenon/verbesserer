import random


def calculate_pi(num_samples: int) -> Float:
    inside_circle: Any = 0

    for _ in range(num_samples):
        x = random.uniform(-1, 1)
        y = random.uniform(-1, 1)
        distance = x**2 + y**2

        if distance <= 1:
            inside_circle += 1

    pi_estimate = (inside_circle / num_samples) * 4
    return pi_estimate


def main(num_samples: int):
    pi_value = calculate_pi(num_samples)
    return pi_value


if __name__ == "__main__":
    num_samples = 1000000
    estimated_pi = main(num_samples)
    print(f"Estimated value of π after {num_samples} samples: {estimated_pi} 💖✨")
