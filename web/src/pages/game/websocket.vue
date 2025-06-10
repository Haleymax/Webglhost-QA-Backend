<template>
  <v-card class="websocket-container mx-auto" max-width="600">
    <v-card-title class="text-h5">处理结果</v-card-title>
    
    <v-card-text>
      <div class="status mb-4">
        连接状态: 
        <v-chip :color="isConnected ? 'success' : 'error'" small>
          {{ isConnected ? '已连接' : '未连接' }}
        </v-chip>
      </div>
      
      <v-textarea
        v-model="receivedMessages"
        label="运行结果"
        readonly
        outlined
        rows="10"
        auto-grow
        style="overflow-y: auto; max-height: 200px;"
        placeholder="接收到的消息将显示在这里..."
        class="mb-4"
      ></v-textarea>
      
      <v-text-field
        v-model="inputMessage"
        label="输入消息"
        outlined
        @keyup.enter="sendMessage"
        placeholder="输入消息..."
        class="mb-2"
      ></v-text-field>
      
      <div class="d-flex flex-wrap gap">
        <v-btn
          color="primary"
          @click="sendMessage"
          :disabled="!isConnected || !inputMessage.trim()"
          class="mr-2 mb-2"
        >
          发送
        </v-btn>
        
        <v-btn
          color="success"
          @click="connectWebSocket"
          :disabled="isConnected"
          class="mr-2 mb-2"
        >
          连接
        </v-btn>
        
        <v-btn
          color="error"
          @click="disconnectWebSocket"
          :disabled="!isConnected"
          class="mr-2 mb-2"
        >
          断开
        </v-btn>

        <v-btn
          color="error"
          @click="handClear"
          >
         清理显示框
        </v-btn>

        
      </div>
    </v-card-text>
  </v-card>
</template>

<script setup lang="ts" name="WebSocket">
import { ref, onMounted, onBeforeUnmount } from 'vue';

const isConnected = ref(false);
const inputMessage = ref('');
const receivedMessages = ref('');
const socket = ref<WebSocket | null>(null);


const wsUrl = 'ws://127.0.0.1:8088/api/v1/game/ws';

const connectWebSocket = () => {
  if (isConnected.value) return;
  
  socket.value = new WebSocket(wsUrl);
  
  socket.value.onopen = () => {
    isConnected.value = true;
    addMessage('系统: WebSocket连接已建立');
  };
  
  socket.value.onmessage = (event) => {
    addMessage(`服务器: ${event.data}`);
  };
  
  socket.value.onerror = (error) => {
    addMessage(`系统: WebSocket错误: ${error}`);
  };
  
  socket.value.onclose = () => {
    isConnected.value = false;
    addMessage('系统: WebSocket连接已关闭');
  };
};

const disconnectWebSocket = () => {
  if (socket.value) {
    socket.value.close();
  }
};

const handClear = () => {
    receivedMessages.value= ""
}

const sendMessage = () => {
  if (!isConnected.value || !inputMessage.value.trim()) return;
  
  try {
    if (socket.value) {
      socket.value.send(inputMessage.value);
      addMessage(`我: ${inputMessage.value}`);
      inputMessage.value = '';
    } else {
      addMessage('系统: WebSocket未连接，无法发送消息');
    }
  } catch (error) {
    addMessage(`系统: 发送消息失败: ${error}`);
  }
};

const addMessage = (message: string) => {
  receivedMessages.value += message + '\n';
};


onMounted(() => {
    connectWebSocket();
});


onBeforeUnmount(() => {
  disconnectWebSocket();
});
</script>