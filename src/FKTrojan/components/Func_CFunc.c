#include <stdio.h>
#include <windows.h>
#include <aclapi.h>
#include <sddl.h>
#include <tchar.h>
#include <Commctrl.h>
#include <stdint.h>
//------------------------------------------------------------
//EnumChildWindows回调函数，hwnd为指定的父窗口
//---------------------------------------------------------
BOOL IsPartOf(char* w1, char* w2)
{
    int i=0;
    int j=0;

    for(i;i < strlen(w1); i++)
    {
        if(w1[i] == w2[j])
        {
            j++;
        }
    }

    if(strlen(w2) == j)
        return TRUE;
    else
        return FALSE;
}
//------------------------------------------------------------
// 最小化Anvir窗口
void MinSizeOfAnvir(){
      HWND baseWnd = NULL;
      for( ; ; ){
          baseWnd = FindWindow("AnVirMainFrame",NULL);
          if(baseWnd == NULL){
              //_tprintf(TEXT("Can't find AnVir window.\n"));
              //return;
          }
          else{
              _tprintf(TEXT("Yes, Found AnVir window.\n"));
              break;
          }
      }
      DWORD dwExStyle = GetWindowLong(baseWnd, GWL_EXSTYLE);
      dwExStyle = WS_EX_TOOLWINDOW;
      SetWindowLong(baseWnd, GWL_EXSTYLE, dwExStyle);

      SetWindowPos(baseWnd,NULL,-200,-200,1,1,SWP_NOZORDER);
      ShowWindow(baseWnd,SW_SHOWNOACTIVATE);
}
//------------------------------------------------------------
// 修改文件最后修改时间
void WriteFileModifyTime(char* filename, long long iModifyTime){

    _tprintf(TEXT("File time: %lld.\n"), iModifyTime);

    FILETIME ft;
    SYSTEMTIME st;
    LONGLONG nLL = Int32x32To64(iModifyTime, 10000000) + 116444736000000000;
    ft.dwLowDateTime = (DWORD)nLL;
    ft.dwHighDateTime = (DWORD)(nLL >> 32);
    FileTimeToSystemTime(&ft, &st);

    char buffer[ 256 ] = {0};
    sprintf( buffer,
             "%d-%02d-%02d %02d:%02d:%02d.%03d",
             st.wYear,
             st.wMonth,
             st.wDay,
             st.wHour,
             st.wMinute,
             st.wSecond,
             st.wMilliseconds );
    _tprintf(TEXT("File time: %s.\n"), buffer);

    FILETIME lastWriteFiletime;
    SystemTimeToFileTime(&st, &lastWriteFiletime);

    // 文件句柄
    HANDLE fileHandler = CreateFile(filename, FILE_WRITE_ATTRIBUTES,
        FILE_SHARE_READ | FILE_SHARE_WRITE, NULL, OPEN_EXISTING,
        FILE_ATTRIBUTE_NORMAL, NULL);

    // 修改文件信息
    SetFileTime(fileHandler, NULL, NULL, &lastWriteFiletime);

    // 关闭
    CloseHandle(fileHandler);
}
//------------------------------------------------------------
// 是的，就是专杀安全狗的
BOOL CALLBACK SetEditText(HWND hWnd,LPARAM lParam)
{
    char WindowClassName[100]={0};
    GetClassName(hWnd,WindowClassName,100);

    if(IsPartOf(WindowClassName, "Edit") && IsWindowVisible(hWnd)){
        HWND hNextWnd = GetNextWindow(hWnd, 2);
        if(hNextWnd != NULL){
            char WindowTextName[100]={0};
            GetWindowText(hNextWnd, WindowTextName, 100);
            if(IsPartOf(WindowTextName, "Name Filter:")){
                    _tprintf(TEXT("Find Edit window: %x.\n"), hWnd);
                    char temp1[32] = {0};
                    strncpy(temp1,"SafeDogGuardCenter",19);
                    SendMessage(hWnd,WM_SETTEXT,0,(LPARAM)temp1);
                    SetWindowText(hNextWnd, "Name filter:");
                    return FALSE;
            }
        }
    }

    return TRUE;
}
//------------------------------------------------------------
BOOL CALLBACK PressSupendBtn(HWND hWnd,LPARAM lParam)
{
    char WindowClassName[100]={0};
    GetClassName(hWnd,WindowClassName,100);

    if(IsPartOf(WindowClassName, "ToolbarWindow32") && IsWindowVisible(hWnd)){
        int vCount = SendMessage(hWnd,TB_BUTTONCOUNT,0,0);
        _tprintf(TEXT("Find Toolbar window: %x - %d.\n"), hWnd, vCount);

        TBBUTTON *_tbn, tbn;
        unsigned long pid;
        HANDLE process;
        GetWindowThreadProcessId(hWnd, &pid);
        process=OpenProcess(PROCESS_VM_OPERATION|PROCESS_VM_READ|PROCESS_VM_WRITE|PROCESS_QUERY_INFORMATION,
         FALSE, pid);
        _tbn=(TBBUTTON*)VirtualAllocEx(process, NULL, sizeof(TBBUTTON), MEM_COMMIT, PAGE_READWRITE);
        SendMessage(hWnd, TB_GETBUTTON, (WPARAM)10, (LPARAM)_tbn);
        TBBUTTON local;
        ReadProcessMemory(process, _tbn, &local, sizeof(TBBUTTON), NULL);
        VirtualFreeEx(process, _tbn, 0, MEM_RELEASE);

        _tprintf(TEXT("Find menuid: %d.\n"), local.idCommand);
        SendMessage(hWnd, WM_COMMAND, local.idCommand, (LPARAM)TRUE);

        return FALSE;
    }
    return TRUE;
}
//------------------------------------------------------------
BOOL CALLBACK SelectFirstResult(HWND hWnd,LPARAM lParam)
{
    char WindowClassName[100]={0};
    GetClassName(hWnd,WindowClassName,100);
    char WindowTextName[100]={0};
    GetWindowText(hWnd, WindowTextName, 100);
    if( IsPartOf(WindowClassName, "SysListView32")
        && IsPartOf(WindowTextName, "AnVirListProcess")
        && IsWindowVisible(hWnd)){

        _tprintf(TEXT("Find Listview window: %x.\n"), hWnd);

        // 强制选中第一个单元
        LVITEM lvi, *_lvi;
        unsigned long pid;
        HANDLE process;
        GetWindowThreadProcessId(hWnd, &pid);
        process=OpenProcess(PROCESS_VM_OPERATION|PROCESS_VM_READ|PROCESS_VM_WRITE|PROCESS_QUERY_INFORMATION,
         FALSE, pid);
        _lvi=(LVITEM*)VirtualAllocEx(process, NULL, sizeof(LVITEM), MEM_COMMIT, PAGE_READWRITE);
        lvi.mask = LVIF_STATE;
        lvi.state = LVIS_SELECTED | LVIS_FOCUSED;
        lvi.stateMask = LVIS_SELECTED | LVIS_FOCUSED;
        WriteProcessMemory(process, _lvi, &lvi, sizeof(LVITEM), NULL);
        SendMessage(hWnd, LVM_SETITEMSTATE, (WPARAM)0, (LPARAM)_lvi);
        VirtualFreeEx(process, _lvi, 0, MEM_RELEASE);

        return FALSE;
    }
    return TRUE;
}
//------------------------------------------------------------
void AutoDoAnvir_SetEditText()
{
    HWND baseWnd = FindWindow("AnVirMainFrame",NULL);
    if(baseWnd == NULL){
        _tprintf(TEXT("Can't find AnVir window.\n"));
        return;
    }
    _tprintf(TEXT("Found AnVir window.\n"));
    // 修改筛选框的文字
    EnumChildWindows(baseWnd,SetEditText,(LPARAM)(NULL));
}
//------------------------------------------------------------
void AutoDoAnvir_SelectFristListItem()
{
    HWND baseWnd = FindWindow("AnVirMainFrame",NULL);
    if(baseWnd == NULL){
        _tprintf(TEXT("Can't find AnVir window.\n"));
        return;
    }
    _tprintf(TEXT("Found AnVir window.\n"));
    // 选中第一单元
    EnumChildWindows(baseWnd,SelectFirstResult,(LPARAM)(NULL));
}
//------------------------------------------------------------
void AutoDoAnvir_ClickBlock()
{
     HWND baseWnd = FindWindow("AnVirMainFrame",NULL);
     if(baseWnd == NULL){
         _tprintf(TEXT("Can't find AnVir window.\n"));
         return;
     }
     _tprintf(TEXT("Found AnVir window.\n"));
     // 按下挂起符
     EnumChildWindows(baseWnd,PressSupendBtn,(LPARAM)(NULL));
}
//------------------------------------------------------------