#ifndef GO_LIBVMI_H
#define GO_LIBVMI_H

addr_t
get_addr_t(unsigned long long val)
{
  addr_t address = val;
  return address;

}

unsigned long long convert_addr_t(addr_t addr)
{
  unsigned long long val;
  val = addr;
  return val;

}
reg_t
get_reg_t(unsigned long long val)
{
  reg_t reg = val;
  return reg;
}
unsigned long long convert_reg_t(reg_t reg)
{
  unsigned long long val;
  val = reg;
  return val;

}
vmi_pid_t
get_vmi_pid_t(int val)
{
  vmi_pid_t pid = val;
  return pid;
}
size_t
get_size_t(unsigned int val)
{
  size_t size = val;
  return size;
}

#endif /* GO_LIBVMI_H */
