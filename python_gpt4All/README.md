# Setup
- Install vulkan https://vulkan.lunarg.com/sdk/home
- Install cmake https://cmake.org/download/
`
git clone --recurse-submodules https://github.com/nomic-ai/gpt4all.git
cd gpt4all/gpt4all-backend/
mkdir build
cd build
cmake ..
cmake --build . --parallel  # optionally append: --config Release
`
